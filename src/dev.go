package main

import (
	"fmt"
	"log"
	zconfig "weixinsdk/src/config"
	zstorage "weixinsdk/src/storage"
	//注意导入Thrift生成的接口包
	//zutils "weixinsdk/src/utils"
	"weixinsdk/src/core/service"
	"weixinsdk/src/logger"
	//"github.com/sirupsen/logrus"
)

func init() {
	//加载全局配置
	err := zconfig.Load()
	if err != nil {
		log.Fatalf("加载配置文件失败:%s", err)
	}

	//加载日志
	err = logger.Load()
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

	//素材列表
	s, err := service.MaterialList()
	if err != nil {
	}
	fmt.Println(s)

	//素材总数
	sb, err := service.MaterialCount()
	if err != nil {
	}
	fmt.Println(sb)

}
