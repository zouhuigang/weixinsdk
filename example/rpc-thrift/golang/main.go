package main

import (
	"weixinsdk/example/rpc-thrift/golang/controller"
	//注意导入Thrift生成的接口包

	"github.com/labstack/echo"
)

//https://blog.csdn.net/liuxinmingcode/article/details/78293146

//https://studygolang.com/articles/3110
func main() {

	e := echo.New()
	//静态文件
	e.Static("/static", "tmp")
	e.HideBanner = true
	frontG := e.Group("")
	controller.RegisterRoutes(frontG)
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
