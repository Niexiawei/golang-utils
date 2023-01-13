package httpresponse

import (
	"encoding/json"
)

type Response struct {
	Code     int32  `json:"code"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Error    any    `json:"error"`
	HttpCode int    `json:"-"`
}

type Context interface {
	JSON(code int, obj any)
}

func NewResponse(code int32, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
	}
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) WithError(err any) *Response {
	r.Error = err
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

func (r *Response) ResultOk(c Context) {
	c.JSON(200, r)
}

func (r *Response) ResultFail(c Context, httpCode ...int) {
	code := 200
	if len(httpCode) >= 1 {
		code = httpCode[0]
	}
	c.JSON(code, r)
}

type ResponseResultOptions func(response *Response)

func ResultWithMsg(msg string) ResponseResultOptions {
	return func(response *Response) {
		response.Message = msg
	}
}

func ResultWithError(error any) ResponseResultOptions {
	return func(response *Response) {
		response.Error = error
	}
}

func ResultWithData(data any) ResponseResultOptions {
	return func(response *Response) {
		response.Data = data
	}
}

func ResultWithHttpCode(code int) ResponseResultOptions {
	return func(response *Response) {
		response.HttpCode = code
	}
}

func Result(c Context, code int32, options ...ResponseResultOptions) {
	response := &Response{
		Code:    code,
		Message: "ok",
	}
	for _, o := range options {
		o(response)
	}
	c.JSON(200, response)
}

func ResultFail(c Context, code int32, options ...ResponseResultOptions) {
	response := &Response{
		Code:     code,
		HttpCode: 200,
	}
	for _, o := range options {
		o(response)
	}
	c.JSON(response.HttpCode, response)
}
