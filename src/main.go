package main

import (
	"fmt"
	"log"

	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	zstorage "weixinsdk/src/storage"
)

func init() {
	//加载全局配置
	err := zconfig.Load("/build/base.env.ini", "/build/dev.env.ini")
	if err != nil {
		log.Fatalf("加载配置文件失败:%s", err)
	}

	//初始化存储引擎
	err = zstorage.Load()
	if err != nil {
		log.Fatalf("存储初始化失败:%s", err)
	}

}

func main() {

	//测试access_token
	//zcache.GetAccessToken()

	//测试jsapi_ticket
	dat := zcache.GetJsapiTicket()

	fmt.Printf("dat:%s\n", dat)
}
