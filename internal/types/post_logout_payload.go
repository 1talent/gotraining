// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostLogoutPayload post logout payload
//
// swagger:model PostLogoutPayload
type PostLogoutPayload struct {

	// Optional refresh token to delete while logging out
	// Example: 700ebed3-40f7-4211-bc83-a89b22b9875e
	// Format: uuid4
	RefreshToken strfmt.UUID4 `json:"refresh_token,omitempty"`
}

// Validate validates this post logout payload
func (m *PostLogoutPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostLogoutPayload) validateRefreshToken(formats strfmt.Registry) error {
	if swag.IsZero(m.RefreshToken) { // not required
		return nil
	}

	if err := validate.FormatOf("refresh_token", "body", "uuid4", m.RefreshToken.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post logout payload based on context it is used
func (m *PostLogoutPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostLogoutPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostLogoutPayload) UnmarshalBinary(b []byte) error {
	var res PostLogoutPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
