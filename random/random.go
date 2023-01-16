package random

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)
import cryptoRand "crypto/rand"

var randomOffsetInt chan int64

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
}

func RandInt(min, max int) int {
	offset := <-randomOffsetInt
	r := rand.New(rand.NewSource(time.Now().UnixNano() + offset))
	return r.Intn(max-min) + min
}

func RandFloat64(max, min float64) float64 {
	offset := <-randomOffsetInt
	r := rand.New(rand.NewSource(time.Now().UnixNano() + offset))
	randomFloat64 := min + r.Float64()*(max-min)
	random, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", randomFloat64), 64)
	return random
}
