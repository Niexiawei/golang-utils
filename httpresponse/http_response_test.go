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
	Result(&responseContext{}, 200, []string{"111", "222"}, "ok")
}

func TestResultFail(t *testing.T) {
	Result(&responseContext{}, 200, []string{"111", "222"}, "", ResultWithError("error"))
}
