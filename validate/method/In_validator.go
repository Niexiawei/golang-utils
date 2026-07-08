package validate_method

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/Niexiawei/golang-utils/slice"
	validate_interface "github.com/Niexiawei/golang-utils/validate/interface"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	_ validate_interface.CustomValidator = (*InValidatorRegister)(nil)
)

// custom_validate
type InValidatorRegister struct {
}

func (i InValidatorRegister) GetTag() string {
	return "in"
}

func (i InValidatorRegister) TransRegister(v *validator.Validate, t *ut.Translator) error {
	return v.RegisterTranslation(i.GetTag(), *t, func(ut ut.Translator) error {
		return ut.Add(i.GetTag(), "{0}不在允许的范围内", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(i.GetTag(), fe.Field())
		return t
	})
}

func (i InValidatorRegister) ValidatorRegister(v *validator.Validate) error {
	return v.RegisterValidation("in", inValidator)
}

func inValidator(fl validator.FieldLevel) bool {
	if fl.Field().IsZero() {
		return true
	}
	param := fl.Param()
	params := strings.Split(param, " ")
	var value string
	switch fl.Field().Kind() {
	case reflect.String:
		value = fl.Field().String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = strconv.FormatInt(fl.Field().Int(), 10)
	default:
		return false
	}
	return slice.Contain(params, value)
}
