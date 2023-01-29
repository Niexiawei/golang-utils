package golangutils

import (
	"reflect"
	"unsafe"
)

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
