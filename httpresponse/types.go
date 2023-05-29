package httpresponse

type Options func(params *ResponseParams)

type ResponseParams struct {
	DefaultErrCode int32
}

func WithDefaultErrCode(code int32) Options {
	return func(params *ResponseParams) {
		params.DefaultErrCode = code
	}
}
