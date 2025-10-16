package error

import (
	"errors"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(status int, message string) *Error {
	return &Error{
		Message: message,
		Status:  status,
	}
}

var (
	ErrInvalidAircraft = New(http.StatusBadRequest, "invalid aircraft type specified")
)

func GetStatus(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status
	}

	var ve *ValidationError
	if errors.As(err, &ve) {
		return ve.Status
	}

	return http.StatusInternalServerError
}
