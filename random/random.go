package random

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
	"unsafe"
)
import cryptoRand "crypto/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

var randomOffsetInt chan int64
var src = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func init() {
	runWriterRandomOffset()
}

func runWriterRandomOffset() {
	randomOffsetInt = make(chan int64, 50)
	go func() {
		for {
			offset := func() int64 {
				n, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(1000))
				return n.Int64()
			}()
			randomOffsetInt <- offset
		}
	}()

	go func() {
		ticker := time.NewTicker(20 * time.Second)
		for {
			select {
			case <-ticker.C:
				offset := <-randomOffsetInt
				src.Seed(time.Now().UnixNano() + offset)
			}
		}
	}()
}

func RandInt(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return src.Intn(max-min) + min
}

func RandFloat64(max, min float64) float64 {
	if min > max {
		min, max = max, min
	}
	randomFloat64 := min + src.Float64()*(max-min)
	random, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", randomFloat64), 64)
	return random
}

func RandStr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
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
