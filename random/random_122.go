//go:build go1.22

package random

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"unsafe"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.IntN(max-min) + min
}

func RandFloat64(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	randomFloat64 := min + rand.Float64()*(max-min)
	random, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", randomFloat64), 64)
	return random
}

func RandStr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, rand.Int64(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int64(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
