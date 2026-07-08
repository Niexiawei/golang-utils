package validate_method

import (
	"fmt"
	"github.com/Niexiawei/golang-utils/slice"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
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
	switch fl.Field().Kind() {
	case reflect.String:
		value = fl.Field().String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = strconv.FormatInt(fl.Field().Int(), 10)
	default:
		return false
	}

	// dive 校验切片/数组元素时 FieldName 形如 "Tags[0]"，取字段本名再拼接方法名
	name := fl.FieldName()
	if idx := strings.IndexByte(name, '['); idx != -1 {
		name = name[:idx]
	}
	funcName := fmt.Sprintf("%sInParams", name)

	parent := fl.Parent()
	method := parent.MethodByName(funcName)
	if !method.IsValid() && parent.CanAddr() {
		// 支持指针接收者定义的 XxxInParams 方法，避免值接收者产生的结构体复制
		method = parent.Addr().MethodByName(funcName)
	}
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
