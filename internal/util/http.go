package util

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/1talent/gotraining/internal/api/httperrors"
	"github.com/1talent/gotraining/internal/types"
	oerrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/labstack/echo/v4"
)

// BindAndValidateBody binds the request, parsing **only** its body (depending on the `Content-Type` request header) and performs validation
// as enforced by the Swagger schema associated with the provided type.
//
// Note: In contrast to BindAndValidate, this method does not restore the body after binding (it's considered consumed).
// Thus use BindAndValidateBody only once per request!
//
// Returns an error that can directly be returned from an echo handler and sent to the client should binding or validating of any model fail.
func BindAndValidateBody(c echo.Context, v runtime.Validatable) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindBody(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

// ValidateAndReturn returns the provided data as a JSON response with the given HTTP status code after performing payload
// validation as enforced by the Swagger schema associated with the provided type.
// `v` must implement `github.com/go-openapi/runtime.Validatable` in order to perform validations, otherwise an internal server error is thrown.
// Returns an error that can directly be returned from an echo handler and sent to the client should sending or validating fail.
func ValidateAndReturn(c echo.Context, code int, v runtime.Validatable) error {
	if err := validatePayload(c, v); err != nil {
		return err
	}

	return c.JSON(code, v)
}

func validatePayload(c echo.Context, v runtime.Validatable) error {
	if err := v.Validate(strfmt.Default); err != nil {

		var compositeError *oerrors.CompositeError
		if errors.As(err, &compositeError) {
			LogFromEchoContext(c).Debug().Errs("validation_errors", compositeError.Errors).Msg("Payload did match schema, returning HTTP validation error")

			valErrs := formatValidationErrors(c.Request().Context(), compositeError)

			return httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)
		}

		var validationError *oerrors.Validation
		if errors.As(err, &validationError) {
			LogFromEchoContext(c).Debug().AnErr("validation_error", validationError).Msg("Payload did match schema, returning HTTP validation error")

			valErrs := []*types.HTTPValidationErrorDetail{
				{
					Key:   &validationError.Name,
					In:    &validationError.In,
					Error: swag.String(validationError.Error()),
				},
			}

			return httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)
		}

		LogFromEchoContext(c).Error().Err(err).Msg("Failed to validate payload, returning generic HTTP error")
		return err
	}

	return nil
}

func formatValidationErrors(ctx context.Context, err *oerrors.CompositeError) []*types.HTTPValidationErrorDetail {
	valErrs := make([]*types.HTTPValidationErrorDetail, 0, len(err.Errors))
	for _, e := range err.Errors {

		var validationError *oerrors.Validation
		if errors.As(e, &validationError) {
			valErrs = append(valErrs, &types.HTTPValidationErrorDetail{
				Key:   &validationError.Name,
				In:    &validationError.In,
				Error: swag.String(validationError.Error()),
			})
			continue
		}

		var compositeError *oerrors.CompositeError
		if errors.As(e, &compositeError) {
			valErrs = append(valErrs, formatValidationErrors(ctx, compositeError)...)
			continue
		}

		LogFromContext(ctx).Warn().Err(e).Str("err_type", fmt.Sprintf("%T", e)).Msg("Received unknown error type while validating payload, skipping")
	}

	return valErrs
}
