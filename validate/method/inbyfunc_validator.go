package validate_method

import (
	"fmt"
	"github.com/Niexiawei/golang-utils/slice"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strconv"
)

const InByFuncValidatorTag = "in_by_func"

type InByFuncRegister struct {
}

func (i InByFuncRegister) GetTag() string {
	return InByFuncValidatorTag
}
func (i InByFuncRegister) TransRegister(v *validator.Validate, t *ut.Translator) error {
	return v.RegisterTranslation(i.GetTag(), *t, func(ut ut.Translator) error {
		return ut.Add(i.GetTag(), "{0}不在允许的范围内", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(i.GetTag(), fe.Field())
		return t
	})
}

func (i InByFuncRegister) ValidatorRegister(v *validator.Validate) error {
	return v.RegisterValidation(InByFuncValidatorTag, inByFuncValidator)
}

func inByFuncValidator(fl validator.FieldLevel) bool {
	var value string
	val := fl.Field().Interface()
	switch val.(type) {
	case int:
		value = strconv.Itoa(val.(int))
		break
	case int64:
		value = strconv.FormatInt(val.(int64), 10)
		break
	default:
		return false
	}
	funcName := fmt.Sprintf("%sInParams", fl.FieldName())
	method := fl.Parent().MethodByName(funcName)
	if !method.IsValid() {
		return false
	}
	callValue := method.Call(nil)[0]
	callValueStringSlice, ok := callValue.Interface().([]string)
	if !ok {
		return false
	}
	return slice.Contain(callValueStringSlice, value)
}
