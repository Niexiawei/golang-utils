package validate

import (
	"errors"
	"github.com/Niexiawei/golang-utils/maputil"
	validate_method "github.com/Niexiawei/golang-utils/validate/method"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chinese "github.com/go-playground/validator/v10/translations/zh"
	"io"
	"reflect"
)

var (
	validatorInstance      *validator.Validate
	validatorTransInstance ut.Translator
)

func Setup(v *validator.Validate) {
	validatorInstance = v
	validatorTransHandler()
	registerCustomFieldName()
	registerAll()
}

func registerAll() {
	validatorMethods := []CustomValidator{
		validate_method.LocationRegister{},
		validate_method.InValidatorRegister{},
		validate_method.InByFuncRegister{},
	}
	for _, v := range validatorMethods {
		v.ValidatorRegister(validatorInstance)
		v.TransRegister(validatorInstance, &validatorTransInstance)
	}
}

func Register(r ...CustomValidator) {
	for _, v := range r {
		v.ValidatorRegister(validatorInstance)
		v.TransRegister(validatorInstance, &validatorTransInstance)
	}
}

type CustomValidator interface {
	TransRegister(v *validator.Validate, t *ut.Translator) error
	ValidatorRegister(v *validator.Validate) error

	GetTag() string
}

func validatorTransHandler() {
	zhCn := zh.New()
	uni := ut.New(zhCn, zhCn)
	validatorTransInstance, _ = uni.GetTranslator("zh")
	_ = chinese.RegisterDefaultTranslations(validatorInstance, validatorTransInstance)
}

func registerCustomFieldName() {
	validatorInstance.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
}

type RequestBindingError map[string]string

func FormatError(err error) *RequestBindingError {
	if errors.Is(err, io.EOF) {
		return &RequestBindingError{
			"errors": "请求体不能为空",
		}
	}

	if err, ok := err.(validator.ValidationErrors); ok {
		_errors := make(RequestBindingError)
		for _, value := range err {
			_errors[value.Field()] = value.Translate(validatorTransInstance)
		}
		return &_errors
	}

	return &RequestBindingError{
		"errors": "字段验证错误",
	}
}

func (req *RequestBindingError) All() map[string]string {
	return *req
}

func (req *RequestBindingError) First() string {
	errorValues := maputil.Values[string, string](*req)
	if len(errorValues) < 1 {
		return ""
	}
	return errorValues[0]
}
