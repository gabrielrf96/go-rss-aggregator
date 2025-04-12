package app

import "net/http"

const DefaultInternalErrorMessage = "An unexpected error has occurred, please try again later"

// An API error that will be returned in the response JSON for the final user.
type APIError struct {
	// Front-facing error message that the client will see.
	Message string
	// A list of suggestions for the client on how to correct the error.
	Corrections []string
	HttpCode    int
	// Internal error message for the app log, which the client will not see.
	LogMessage string
}

func (e *APIError) Error() string {
	return e.Message
}

func NewInternalError(logMsg string) *APIError {
	return &APIError{
		Message:    DefaultInternalErrorMessage,
		HttpCode:   http.StatusInternalServerError,
		LogMessage: logMsg,
	}
}

func NewValidationError(msgs []string) *APIError {
	return &APIError{
		Message:     "Request parameter validation failed, see suggested corrections",
		HttpCode:    http.StatusUnprocessableEntity,
		Corrections: msgs,
	}
}

func NewValidationErrorFromString(msg string) *APIError {
	return NewValidationError([]string{msg})
}

func NewConflictError(msg string) *APIError {
	return &APIError{
		Message:  msg,
		HttpCode: http.StatusConflict,
	}
}

func NewNotFoundError(msg string) *APIError {
	return &APIError{
		Message:  msg,
		HttpCode: http.StatusUnprocessableEntity,
	}
}
