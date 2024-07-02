package middlewares

import (
	"github.com/Niexiawei/golang-utils/condition"
	"github.com/gofrs/uuid/v5"
	"net/http"
)

type ginContext interface {
	Header(key, value string)
	Next()
	AbortWithStatus(code int)
	GetHeader(key string) string
}

func Cors[T ginContext](c T, method string) {
	origin := c.GetHeader("Origin")
	c.Header("Access-Control-Allow-Origin", condition.If(origin == "", "*", origin))
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token,zgf-user-center-token")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Private-Network", "true")
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}

type GinMiddle struct {
	context ginContext
}

func (g *GinMiddle) Cors(method string) {
	origin := g.context.GetHeader("Origin")
	g.context.Header("Access-Control-Allow-Origin", condition.If(origin == "", "*", origin))
	g.context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	g.context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token,zgf-user-center-token")
	g.context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	g.context.Header("Access-Control-Allow-Credentials", "true")
	g.context.Header("Access-Control-Allow-Private-Network", "true")
	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		g.context.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	g.context.Next()
}

func RequestId[T ginContext]() func(T) {
	return func(t T) {
		u1, _ := uuid.NewV1()
		u := uuid.NewV3(u1, u1.String())
		t.Header("Request-Id", u.String())
		t.Next()
	}
}
