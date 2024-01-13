package ginmiddlewares

import (
	"github.com/Niexiawei/golang-utils/condition"
	"net/http"
)

type ginContext interface {
	Header(key, value string)
	Next()
	AbortWithStatus(code int)
}

func Cors[T ginContext](method, origin string) func(T) {
	return func(c T) {
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
}
