package golangutils

import "testing"

func TestChunkString(t *testing.T) {
	t.Log(ChunkString("测试测试测试", 3))
}
