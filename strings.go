package golangutils

import "unsafe"

// Deprecated: 移动到 stringstools 包内
func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

// Deprecated: 移动到 stringstools 包内
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// Deprecated: 移动到 stringstools 包内
func ChunkString[T ~string](str T, size int) []T {
	if size <= 0 {
		panic("Size parameter must be greater than 0")
	}

	if len(str) == 0 {
		return []T{""}
	}

	if size >= len(str) {
		return []T{str}
	}

	var (
		chunks = []T{}
	)

	chunks = make([]T, 0, ((len(str)-1)/size)+1)
	currentLen := 0
	currentStart := 0
	for i := range str {
		if currentLen == size {
			chunks = append(chunks, str[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, str[currentStart:])
	return chunks
}
