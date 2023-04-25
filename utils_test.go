package golangutils

import (
	"fmt"
	"testing"
)

func TestGetStringRuntimeStack(t *testing.T) {
	fmt.Println(GetStringRuntimeStack(0))
}

func TestBytesToString(t *testing.T) {
	t.Log(BytesToString([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}))
}
