package helper

import (
	"math/rand"
	"time"
)

var HelperFunc = helperFunc{}

type helperFunc struct{}

func (*helperFunc)RandString(len int) string {
	rr := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rr.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}