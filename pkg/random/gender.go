package random

import (
	"math/rand"
	"time"
)

//随机生成0 男、1 女
func GenGender() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2)
}
