package golangutils

import (
	"testing"

	"github.com/Niexiawei/golang-utils/strings"
)

func TestBytesToString(t *testing.T) {
	t.Log(strings.BytesToString([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}))
}

func TestStringToBytes(t *testing.T) {
	t.Log(strings.StringToBytes("test string to bytes"))
}
