package main

import (
	"fmt"
	"log"

	zconfig "weixinsdk/src/config"
	zstorage "weixinsdk/src/storage"
)

func init() {
	//加载全局配置
	err := zconfig.Load("/build/base.env.ini", "/build/dev.env.ini")
	if err != nil {
		log.Fatalf("加载配置文件失败:%s", err)
	}
}

func main() {

	storage_type := zconfig.CFG.MustValue("parameter", "storage_type", "local")
	exporter := zstorage.ExporterMap()[storage_type]

	var ok bool
	var storageInterface zstorage.Exporter
	storageInterface, ok = exporter.(zstorage.Exporter)
	if !ok {
		log.Fatalf("storage init fail")
		return
	}
	err := storageInterface.New()
	if err != nil {
		log.Fatalf("storage new fail:%s", err)
		return
	}

	fmt.Println("==========", storageInterface.Get("weixin_dev"), ok)

	fmt.Println("success", storage_type)
}
