package util

import "github.com/go-playground/validator/v10"

type XValidator struct {
	Validator *validator.Validate
}

type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
}

func (v *XValidator) Validate(data interface{}) []ValidationErrorResponse {
	errors := make([]ValidationErrorResponse, 0)

	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationError := &ValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			}

			errors = append(errors, *validationError)
		}
	}
	return errors
}
