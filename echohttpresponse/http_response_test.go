package echohttpresponse

import (
	"encoding/json"
	"fmt"
	"testing"
)

type responseContext struct {
}

func (r *responseContext) JSON(code int, obj any) error {
	jj, _ := json.Marshal(obj)
	fmt.Printf("code:%d;\ndata:%+v", code, string(jj))
	return nil
}

func TestResult(t *testing.T) {
	Result(&responseContext{}, 200, []string{"111", "222"})
}

func TestResultFail(t *testing.T) {
	ResultFail(&responseContext{}, 200, "error", ResultWithError("error"), ResultWithData([]string{"hahahha", "xixixi"}))
}
