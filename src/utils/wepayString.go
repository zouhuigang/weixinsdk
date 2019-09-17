package utils

import (
	"encoding/xml"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//生成随机字符
func GenerateNonceString() string {
	dict := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var randStr string
	for i := 0; i < 32; i++ {
		index := rand.Intn(35)
		randStr += dict[index : index+1]
	}

	return randStr
}

//生成验证字符
func GenerateSignString(data map[string]interface{}, key string) (str string, err error) {

	delete(data, "sign")

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var paramsStr string
	for _, k := range keys {
		if k == "key" {
			continue
		}
		if reflect.TypeOf(data[k]).Kind() == reflect.Int {
			if data[k].(int) == 0 {
				continue
			}
			paramsStr += k + "=" + strconv.Itoa(data[k].(int)) + "&"
		} else if reflect.TypeOf(data[k]).Kind() == reflect.String {
			if data[k].(string) == "" {
				continue
			}
			paramsStr += k + "=" + data[k].(string) + "&"
		}

	}

	paramsStr = paramsStr + "key=" + key

	paramsStr = MD5(paramsStr)

	paramsStr = strings.ToUpper(paramsStr)

	return paramsStr, nil
}

//生成XML
func GenerateRequestXml(params interface{}) (str string, err error) {

	data, err := xml.MarshalIndent(&params, "", " ")

	if err != nil {
		return "", nil
	}

	return string(data), nil
}

//驼峰转下划线
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
