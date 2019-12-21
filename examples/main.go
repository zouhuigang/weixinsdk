package main

import (
	"log"
	"weixinsdk/examples/controller"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	zstorage "weixinsdk/src/storage"

	"github.com/labstack/echo"
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
	e := echo.New()
	//静态文件
	e.Static("/static", "tmp")
	e.HideBanner = true
	frontG := e.Group("")
	controller.RegisterRoutes(frontG)
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
