package utils

import (
	"bytes"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/zouhuigang/package/zcrypto"
)

//NonceStr:   必须, 32个字符以内, 商户生成的随机字符串
//随机32位字符串
func NonceStr() string {
	chars := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UnixNano())
	result := [32]byte{}
	for i := 0; i < 32; i++ {
		result[i] = chars[rand.Int31n(35)]
	}
	return string(result[:])
}

//过滤
func paraFilter(params map[string]string) []string {
	keys := make([]string, 0)
	for k, v := range params {
		if k == "sign" || strings.TrimSpace(v) == "" {
			delete(params, k)
		} else {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}

//拼接字符串 按照“参数=参数值”的模式用“&”字符拼接成字符串
func createLinkString(keys []string, args map[string]string) string {
	buf := bytes.NewBufferString("")
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(args[k])
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

//签名,微信调试签名https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign
//微信的签名函数sha1(string1)是php版的,可能会跟其他语言不一样
func Signature(args map[string]string) string {
	keys := paraFilter(args)
	signStr := createLinkString(keys, args)
	sign := zcrypto.PhpSha1(signStr)
	return sign
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
