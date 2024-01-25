package validate_method

import (
	"github.com/Niexiawei/golang-utils/slice"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

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
	var value string
	if fl.Field().IsZero() {
		return true
	}
	val := fl.Field().Interface()
	param := fl.Param()
	params := strings.Split(param, " ")
	switch val.(type) {
	case int:
		value = strconv.Itoa(val.(int))
		break
	case int64:
		value = strconv.FormatInt(val.(int64), 10)
		break
	case string:
		value = val.(string)
	default:
		return false
	}
	return slice.Contain(params, value)
}
