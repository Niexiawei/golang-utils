package httpresponse

import (
	"encoding/json"
)

type Response struct {
	Code           int32   `json:"code"`
	Message        string  `json:"message,omitempty"`
	Data           any     `json:"data,omitempty"`
	Error          any     `json:"error,omitempty"`
	HttpCode       int     `json:"-"`
	Context        Context `json:"-"`
	defaultErrCode int32
}

type Context interface {
	JSON(code int, obj any)
}

func NewDefaultResponse(c Context, option ...Options) *Response {
	params := &ResponseParams{
		DefaultErrCode: 500,
	}
	for _, o := range option {
		o(params)
	}
	return &Response{
		Context:        c,
		Message:        "ok",
		Code:           200,
		defaultErrCode: params.DefaultErrCode,
	}
}

func NewResponse(c Context, code int32, msg string, option ...Options) *Response {
	params := &ResponseParams{
		DefaultErrCode: 500,
	}
	for _, o := range option {
		o(params)
	}
	return &Response{
		Code:    code,
		Message: msg,
		Context: c,
	}
}

func (r *Response) WithCode(code int32) *Response {
	r.Code = code
	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) WithError(err any) *Response {
	r.Error = err
	r.Code = r.defaultErrCode
	r.Message = ""
	return r
}

func (r *Response) WithErrorAndCode(err any, code int32) *Response {
	r.Error = err
	r.Code = code
	return r
}

func (r *Response) WithMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) ToString() string {
	byteStr, _ := json.Marshal(*r)
	return string(byteStr)
}

func (r *Response) Get() Response {
	return *r
}

func (r *Response) ResultOk() {
	r.Context.JSON(200, r)
}

func (r *Response) ResultFail(httpCode ...int) {
	code := 200
	if len(httpCode) >= 1 {
		code = httpCode[0]
	}
	r.Context.JSON(code, r)
}

type ResponseResultOptions func(response *Response)

func ResultWithError(error any) ResponseResultOptions {
	return func(response *Response) {
		response.Error = error
	}
}

func ResultWithHttpCode(code int) ResponseResultOptions {
	return func(response *Response) {
		response.HttpCode = code
	}
}

func Result(c Context, code int32, data any, msg string, options ...ResponseResultOptions) {
	response := &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	for _, o := range options {
		o(response)
	}
	c.JSON(200, response)
}

func ResultFail(c Context, code int32, data any, msg string, options ...ResponseResultOptions) {
	response := &Response{
		Code:     code,
		HttpCode: 200,
		Message:  msg,
		Data:     data,
	}
	for _, o := range options {
		o(response)
	}
	c.JSON(response.HttpCode, response)
}
