package httpresponse

import "fmt"

type Response struct {
	Code           int32   `json:"code"`
	Message        string  `json:"message,omitempty"`
	Data           any     `json:"data,omitempty"`
	Error          any     `json:"error,omitempty"`
	HttpCode       int     `json:"-"`
	Context        Context `json:"-"`
	defaultErrCode int32
}

type ResponseResultOptions func(response *Response)

func ResultWithError(error any) ResponseResultOptions {
	return func(response *Response) {
		response.Error = anyToString(error)
	}
}

func ResultWithHttpCode(code int) ResponseResultOptions {
	return func(response *Response) {
		response.HttpCode = code
	}
}

func ResultWithData(data any) ResponseResultOptions {
	return func(response *Response) {
		response.Data = data
	}
}

func ResultWithMsg(msg any) ResponseResultOptions {
	return func(response *Response) {
		response.Message = anyToString(msg)
	}
}

func Result(c Context, code int32, data any, options ...ResponseResultOptions) {
	response := &Response{
		Code:    code,
		Message: "ok",
		Data:    data,
	}
	for _, o := range options {
		o(response)
	}
	c.JSON(200, response)
}

func ResultFail(c Context, code int32, msg any, options ...ResponseResultOptions) {
	response := &Response{
		Code:     code,
		HttpCode: 200,
		Message:  anyToString(msg),
		Data:     nil,
	}
	for _, o := range options {
		o(response)
	}
	c.JSON(response.HttpCode, response)
}

func anyToString(msg any) string {
	Msg := ""
	switch msg.(type) {
	case error:
		Msg = msg.(error).Error()
		break
	case string:
		Msg = msg.(string)
		break
	default:
		Msg = fmt.Sprintf("%v", msg)
	}
	return Msg
}
