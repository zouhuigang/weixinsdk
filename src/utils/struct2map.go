package utils

import (
	"encoding/json"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
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
