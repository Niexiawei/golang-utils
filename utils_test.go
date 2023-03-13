package golangutils

import (
	"fmt"
	"testing"
)

func TestGetStringRuntimeStack(t *testing.T) {
	fmt.Println(GetStringRuntimeStack(0))
}
