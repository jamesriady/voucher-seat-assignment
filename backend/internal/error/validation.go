package error

import "net/http"

type ValidationError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
	Status  int               `json:"-"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(errors map[string]string) *ValidationError {
	return &ValidationError{
		Message: "Input validation failed",
		Errors:  errors,
		Status:  http.StatusBadRequest,
	}
}
