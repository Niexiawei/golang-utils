package echohttpresponse

type Context interface {
	JSON(code int, i interface{}) error
}
