package md5lib

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 生成MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// Md5 生成MD5
func Md5(buf []byte)  string{
	dataMd5 := md5.Sum(buf)
	return hex.EncodeToString(dataMd5[:])
}
