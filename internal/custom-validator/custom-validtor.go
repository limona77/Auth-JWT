package custom_validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		Validator *validator.Validate
	}
)

var Validate = validator.New()

func (xV *XValidator) Validate(data interface{}) error {
	err := xV.Validator.Struct(data)
	if err != nil {
		fieldErr := err.(validator.ValidationErrors)[0]

		return xV.wrapValidationError(fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
	}

	return nil
}

func (xV *XValidator) wrapValidationError(field string, tag string, param string) error {
	switch tag {
	case "required":
		return fmt.Errorf("поле %s обязательно для заполнения", field)
	case "email":
		return fmt.Errorf("поле %s должно быть электронной почтой", field)
	case "min":
		return fmt.Errorf("поле %s должно содержать как минимум %s символов", field, param)
	default:
		return fmt.Errorf("поле %s некорректно", field)
	}
}
