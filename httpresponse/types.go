package httpresponse

type Context interface {
	JSON(code int, obj any)
}
