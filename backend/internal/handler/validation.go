package handler

import (
	"fmt"

	internalError "backend/internal/error"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(validate *validator.Validate, s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	errs := make(map[string]string)
	for _, e := range validationErrors {
		jsonFieldName := e.Field()

		switch e.Tag() {
		case "required":
			errs[jsonFieldName] = fmt.Sprintf("The %s field is required.", jsonFieldName)
		case "datetime":
			errs[jsonFieldName] = fmt.Sprintf("The %s field must be in YYYY-MM-DD format.", jsonFieldName)
		default:
			// A fallback message for any other validation rules.
			errs[jsonFieldName] = fmt.Sprintf("The %s field is invalid.", jsonFieldName)
		}
	}

	return internalError.NewValidationError(errs)
}
