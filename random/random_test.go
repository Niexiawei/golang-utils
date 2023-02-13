package random

import (
	"fmt"
	"testing"
)

func Test_int_swap(t *testing.T) {
	min, max := 10, 20
	min, max = max, min
	fmt.Println(min, max)
}

func Test_int_swap_point(t *testing.T) {
	swap := func(a, b *int) {
		*a, *b = *b, *a
	}
	min, max := 10, 20
	swap(&min, &max)
	fmt.Println(min, max)
}

func Test_RandomStr(t *testing.T) {
	t.Log(RandStr(32))
}
