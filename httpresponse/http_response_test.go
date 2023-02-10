package httpresponse

import (
	"fmt"
	"testing"
)

type responseContext struct {
}

func (r *responseContext) JSON(code int, obj any) {
	fmt.Printf("code:%d;\ndata:%+v", code, obj)
}

func TestResult(t *testing.T) {
	c := &responseContext{}
	Result(c, 200, ResultWithData(map[string]string{
		"xixi": "6666",
	}))
	fmt.Printf("\n")
	ResultFail(c, 500, ResultWithMsg("用户认证失败"), ResultWithHttpCode(401))
}

func TestNewResponse(t *testing.T) {
	c := &responseContext{}
	NewResponse(c, 200, "ok").WithData(
		map[string]string{
			"xixi": "6666",
		}).ResultOk()

	NewResponse(c, 500, "ok").WithData(
		map[string]string{
			"xixi": "6666",
		}).WithError("页面不存在").WithMessage("页面不存在！").ResultFail(404)
}
