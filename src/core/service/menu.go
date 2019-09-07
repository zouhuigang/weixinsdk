package service

import (
	"encoding/json"
	"fmt"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service"

	"github.com/zouhuigang/package/zhttp"
)

// 自定义菜单创建接口
func (this *WxServiceThrift) CreateMenu(menu *z_weixin_service.Menu) (*z_weixin_service.WxResponse, error) {

	menuJson, err := json.Marshal(menu)
	if err != nil {
		logger.MyLogger.Error("CreateMenu" + err.Error())
		return nil, err
	}
	response, err := this.CreateMenuByJson(menuJson)
	return response, err
}

// 自定义菜单创建接口（前端定义的菜单json串，而不是在代码内部固定的）
func (this *WxServiceThrift) CreateMenuByJson(menuJsonBytes []byte) (*z_weixin_service.WxResponse, error) {
	//menuJsonByte := bytes.NewReader(menuJsonBytes)
	logger.MyLogger.Debug("CreateMenuByJson:" + string(menuJsonBytes))
	response := z_weixin_service.NewWxResponse()
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}

	requrl := fmt.Sprintf("%s?access_token=%s", zconfig.SERVICE_APIURL_MENU_CREATE, access_token)

	err = zhttp.POSTWithUnmarshal(requrl, menuJsonBytes, response)
	return response, err
}
