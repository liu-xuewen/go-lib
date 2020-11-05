package randomlib

import (
	"math/rand"
	"time"
)


var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")
var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=_")

func init() {
	rand.Seed(time.Now().Unix())
}

// RandStringLow 随机字符串，包含 1~9 和 a~z - [i,l,o]
func RandStringLow(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 31
		b[i] = letters[arc]
	}
	return b
}

// RandString 随机字符串，包含 英文字母和数字附加=_两个符号
func RandString(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	for i, x := range b {
		arc = x & 63
		b[i] = longLetters[arc]
	}
	return string(b)
}