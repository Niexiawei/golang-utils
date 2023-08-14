package golangutils

import (
	"testing"
)

func TestBytesToString(t *testing.T) {
	t.Log(BytesToString([]byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}))
}

func TestStringToBytes(t *testing.T) {
	t.Log(StringToBytes("test string to bytes"))
}
