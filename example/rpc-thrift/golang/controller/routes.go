package controller

import (
	"weixinsdk/example/rpc-thrift/golang/controller/api_1_0"

	"github.com/labstack/echo"
)

func RegisterRoutes(g *echo.Group) {
	//api
	new(api_1_0.DocApi).RegisterRoute(g)
	new(api_1_0.WeixinApi).RegisterRoute(g)
}
