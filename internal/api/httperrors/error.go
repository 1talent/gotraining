package httperrors

import (
	"fmt"
	"sort"
	"strings"

	"github.com/1talent/gotraining/internal/types"
)

const (
	// HTTPErrorTypeGeneric represents the generic error type returned as default for all HTTP errors without a type defined.
	HTTPErrorTypeGeneric string = "generic"
)

// Payload in accordance with RFC 7807 (Problem Details for HTTP APIs) with the exception of the type
// value not being represented by a URI. https://tools.ietf.org/html/rfc7807 @ 2020-04-27T15:44:37Z

type HTTPError struct {
	types.PublicHTTPError
	Internal       error                  `json:"-"`
	AdditionalData map[string]interface{} `json:"-"`
}

type HTTPValidationError struct {
	types.PublicHTTPValidationError
	Internal       error                  `json:"-"`
	AdditionalData map[string]interface{} `json:"-"`
}

func (e *HTTPError) Error() string {
	var b strings.Builder

	fmt.Fprintf(&b, "HTTPError %d (%s): %s", *e.Code, *e.Type, *e.Title)

	if len(e.Detail) > 0 {
		fmt.Fprintf(&b, " - %s", e.Detail)
	}
	if e.Internal != nil {
		fmt.Fprintf(&b, ", %v", e.Internal)
	}
	if e.AdditionalData != nil && len(e.AdditionalData) > 0 {
		keys := make([]string, 0, len(e.AdditionalData))
		for k := range e.AdditionalData {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		b.WriteString(". Additional: ")
		for i, k := range keys {
			fmt.Fprintf(&b, "%s=%v", k, e.AdditionalData[k])
			if i < len(keys)-1 {
				b.WriteString(", ")
			}
		}
	}

	return b.String()
}

func (e *HTTPValidationError) Error() string {
	var b strings.Builder

	fmt.Fprintf(&b, "HTTPValidationError %d (%s): %s", *e.Code, *e.Type, *e.Title)

	if len(e.Detail) > 0 {
		fmt.Fprintf(&b, " - %s", e.Detail)
	}
	if e.Internal != nil {
		fmt.Fprintf(&b, ", %v", e.Internal)
	}
	if e.AdditionalData != nil && len(e.AdditionalData) > 0 {
		keys := make([]string, 0, len(e.AdditionalData))
		for k := range e.AdditionalData {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		b.WriteString(". Additional: ")
		for i, k := range keys {
			fmt.Fprintf(&b, "%s=%v", k, e.AdditionalData[k])
			if i < len(keys)-1 {
				b.WriteString(", ")
			}
		}
	}

	b.WriteString(" - Validation: ")
	for i, ve := range e.ValidationErrors {
		fmt.Fprintf(&b, "%s (in %s): %s", *ve.Key, *ve.In, *ve.Error)
		if i < len(e.ValidationErrors)-1 {
			b.WriteString(", ")
		}
	}

	return b.String()
}
