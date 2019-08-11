package api_1_0

import (
	"weixinsdk/example/rpc-thrift/golang/global"

	"github.com/labstack/echo"
)

type DocApi struct{}

var docApi = DocApi{}

//注册路由
func (this *DocApi) RegisterRoute(g *echo.Group) {
	g.POST("/api/v1.0/doc/list", this.doc_list)
}

func (this *DocApi) doc_list(ctx echo.Context) error {
	data := map[string]interface{}{}

	return global.ResponeJson(ctx, global.ErrOk, data)
}
