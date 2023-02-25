package utils

import (
	"math/rand"
	"time"
)

func GetNumber() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Int63n(10000000)
}
