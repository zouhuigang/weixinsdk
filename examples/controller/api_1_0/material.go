package api_1_0

import (
	"weixinsdk/examples/system"
	zcore "weixinsdk/src/core/service"

	"github.com/labstack/echo"
)

type MaterialApi struct{}

var materialApi = MaterialApi{}

func (this *MaterialApi) RegisterRoute(g *echo.Group) {
	g.Any("/material/list", this.materialList)
}

func (this *MaterialApi) materialList(ctx echo.Context) error {
	data := map[string]interface{}{}

	handler := &zcore.WxServiceThrift{}
	resCoun, _ := handler.MaterialCount()
	list, _ := handler.MaterialList("image", 0, 15)

	data["count"] = resCoun
	data["list"] = list
	return system.ResponeJson(ctx, system.ErrOk, data)
}
