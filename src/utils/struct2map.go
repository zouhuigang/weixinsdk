package utils

import (
	"encoding/json"
	"reflect"
)

//结构体转MAP，支持指针类型
/*
 map[
  appid:wxadd472a86212b893
  body:测试商品
  nonce_str:62ME6IA5BF46WO84LK7QAG3DMHMDPUL3
  notify_u_r_l:http://notify o
  penid:obxjktz34If5J6xal0HXGnK23H8E
  out_trade_no:C20190214000133
  spbill_create_i_p:127.0.0.1
  total_fee:1 trade_type:JSAPI
  ]
*/
func Struct2Map(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	var new_type reflect.Type
	var new_value reflect.Value

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if k := t.Kind(); k == reflect.Ptr { //结构体指针
		new_value = v.Elem()
		new_type = new_value.Type()
	} else {
		new_type = t
		new_value = v
	}

	for i := 0; i < new_type.NumField(); i++ {
		field := new_value.Field(i)
		field_t := new_type.Field(i)
		if field.Kind() == reflect.Struct {
			continue
		} else if field.Kind() == reflect.String && field.Interface().(string) == "" {
			continue
		} else if field.Kind() == reflect.Int && field.Interface().(int) == 0 {
			continue
		} else {
			key := SnakeString(field_t.Name)
			data[key] = field.Interface()
		}
	}
	return data
}

//struct 2 map
func Struct2MapString(obj interface{}) map[string]string {
	m := map[string]string{}
	j, _ := json.Marshal(obj)
	json.Unmarshal(j, &m)

	return m
}
