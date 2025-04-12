package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/RussellLuo/validating/v3"
	"github.com/gabrielrf96/go-rss-aggregator/internal/app"
)

const (
	ValidationRequired = "is required"
	ValidationURL      = "must be a valid URL"
)

type RequestParams interface {
	comparable
	Schema() validating.Schema
}

// Parses the JSON params from the provided [http.Request] into a struct of the provided type [T].
func ParseParams[T RequestParams](r *http.Request) (T, error) {
	var params T

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil && !errors.Is(err, io.EOF) {
		return params, &app.APIError{
			HttpCode:   http.StatusBadRequest,
			Message:    "Error parsing JSON",
			LogMessage: fmt.Sprintf("Error parsing JSON: %v", err),
		}
	}

	var empty T
	if params == empty {
		return params, &app.APIError{
			HttpCode: http.StatusBadRequest,
			Message:  "Empty request body, missing parameters",
		}
	}

	return params, nil
}

// Validates the provided request params based on their defined [validating.Schema].
// If validation fails an [app.APIError] is returned with the list of validation error messages.
func ValidateParams[T RequestParams](params T) error {
	schema := params.Schema()
	if len(schema) == 0 {
		return nil
	}

	errs := validating.Validate(schema)
	errCount := len(errs)

	if errCount == 0 {
		return nil
	}

	apiErr := app.NewValidationError(make([]string, 0, errCount))

	for _, err := range errs {
		apiErr.Corrections = append(
			apiErr.Corrections,
			fmt.Sprintf("Parameter '%s' %s", err.Field(), err.Message()),
		)
	}

	return apiErr
}
