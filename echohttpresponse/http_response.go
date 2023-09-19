package echohttpresponse

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
		response.Error = error
	}
}

func ResultWithHttpCode(code int) ResponseResultOptions {
	return func(response *Response) {
		response.HttpCode = code
	}
}

func Result(c Context, code int32, data any, msg string, options ...ResponseResultOptions) error {
	response := &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	for _, o := range options {
		o(response)
	}
	return c.JSON(200, response)
}

func ResultFail(c Context, code int32, data any, msg string, options ...ResponseResultOptions) error {
	response := &Response{
		Code:     code,
		HttpCode: 200,
		Message:  msg,
		Data:     data,
	}
	for _, o := range options {
		o(response)
	}
	return c.JSON(response.HttpCode, response)
}
