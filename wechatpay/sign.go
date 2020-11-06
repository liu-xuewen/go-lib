package wechatpay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strings"
)


// SecretKey key为商户平台设置的密钥key
var SecretKey string

// SetSecretKey set secretKey
func SetSecretKey(secretKey string) {
	SecretKey = secretKey
}


// Sign 签名
func Sign(paramMap map[string]string) string {
	// 创建切片
	var keys = make([]string, 0, len(paramMap))
	// 遍历签名参数
	for k := range paramMap {
		if k != "sign"|| k!="" { // 排除sign字段,集合M内非空参数值
			keys = append(keys, k)
		}
	}

	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		if len(paramMap) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(paramMap[k])
			buf.WriteString(`&`)
		}
	}
	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(SecretKey)

	var (
		dataMd5    [16]byte
		str        string
	)
	dataMd5 = md5.Sum(buf.Bytes())
	str = hex.EncodeToString(dataMd5[:])
	return strings.ToUpper(str)
}


// VerifySign 校验签名
func VerifySign(req interface{}, sign string) (bool,error) {
	reqByte, err := json.Marshal(req)
	if err != nil {
		return false,err
	}

	var paramMap map[string]string
	err = json.Unmarshal(reqByte, &paramMap)
	if err != nil {
		return false,err
	}
	return sign == Sign(paramMap),nil
}


