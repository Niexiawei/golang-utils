//go:build go1.20

package golangutils

import "unsafe"

func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
