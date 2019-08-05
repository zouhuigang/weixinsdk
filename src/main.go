package main

import (
	"fmt"
	"log"

	zconfig "weixinsdk/src/config"
)

func init() {
	//加载全局配置
	err := zconfig.Load("/build/base.env.ini", "/build/dev.env.ini")
	if err != nil {
		log.Fatalf("加载配置文件失败:%s", err)
	}
}

func main() {

	value, err := zconfig.CFG.GetValue("SERVICE", "TEST")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_default", err)
	}
	fmt.Println(value)
}
