package random

import (
	"math/rand"
	"time"
)

// 生产n位数字，主用作验证码
func RandomInt(length int) (str string) {

	var arr []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := len(arr)
	for i := 0; i < length; i++ {
		str += string(arr[r.Intn(size)])
	}
	return
}
