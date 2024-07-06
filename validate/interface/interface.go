package validate_interface

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type CustomValidator interface {
	TransRegister(v *validator.Validate, t *ut.Translator) error
	ValidatorRegister(v *validator.Validate) error

	GetTag() string
}
