package middlewares

import (
	"fmt"
	"testing"
)

type testCtx struct {
}

func (t *testCtx) Header(key, value string) {
	fmt.Println("set header:", key, value)
}

func (t *testCtx) Next() {
	fmt.Println("next")
}

func (t *testCtx) AbortWithStatus(code int) {
	fmt.Println("abort code:", code)
}

func (t *testCtx) GetHeader(key string) string {
	return ""
}

func TestGinMiddle_Cors(t *testing.T) {
	g := NewGinMiddle(&testCtx{})
	g.Cors("OPTIONS")
}
