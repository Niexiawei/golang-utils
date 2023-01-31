package golangutils

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

var (
	AAAErr = errors.New("哈哈哈哈")
)

func Test_Stack(t *testing.T) {
	err := errors.New("哈哈哈哈")
	stackErr := NewStackErr(AAAErr)
	s := fmt.Errorf("%v\n%w", err, stackErr)
	fmt.Println(s)
	fmt.Printf("%+v", stackErr)
}

func Test_call(t *testing.T) {
	pc, _, _, _ := runtime.Caller(0)
	pc2 := uintptr(pc) - 1
	info := runtime.FuncForPC(pc2)
	fmt.Println(info.Name())
	fmt.Println(info.FileLine(pc2))
}
