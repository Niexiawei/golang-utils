package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	res := Result(200, []string{"111", "222"})
	resB, _ := json.Marshal(res)
	fmt.Println(string(resB))
}

func TestResultFail(t *testing.T) {
	err := errors.New("哈哈哈ERR")
	res := ResultFail(200, err,
		ResultWithError(map[string]string{"aa": "aa", "bb": "bb"}),
		ResultWithData([]string{"hahahha", "xixixi"}),
	)
	resB, _ := json.Marshal(res)
	fmt.Println(string(resB))
}
