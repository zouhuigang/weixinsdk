package controller

import (
	"weixinsdk/examples/controller/api_1_0"

	"github.com/labstack/echo"
)

func RegisterRoutes(g *echo.Group) {
	//api
	new(api_1_0.WeixinApi).RegisterRoute(g)
	new(api_1_0.MaterialApi).RegisterRoute(g)
}
