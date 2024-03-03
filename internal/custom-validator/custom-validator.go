package custom_validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		validator *validator.Validate
	}
)

var validate = validator.New()

func (xV *XValidator) Validate(data interface{}) error {

	err := validate.Struct(data)
	if err != nil {
		fieldErr := err.(validator.ValidationErrors)[0]

		return xV.wrapValidationError(fieldErr.Field(), fieldErr.Tag())
	}

	return nil
}
func (xV *XValidator) wrapValidationError(field string, tag string) error {
	switch tag {
	case "required":
		return fmt.Errorf("field %s is required", field)
	case "email":
		return fmt.Errorf("field %s must be a valid email address", field)
	default:
		return fmt.Errorf("field %s is invalid", field)
	}
}
