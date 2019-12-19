package service

import (
	"encoding/json"
	"fmt"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service"

	"github.com/zouhuigang/package/zhttp"
)

func MaterialCount() (*z_weixin_service.MaCount, error) {
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}

	mc := new(z_weixin_service.MaCount)

	requrl := fmt.Sprintf("%s?access_token=%s", zconfig.SERVICE_APIURL_MATERIAL_COUNT, access_token)

	err = zhttp.GETWithUnmarshal(requrl, &mc)
	return mc, err
}

/*素材列表*/
func MaterialList() (*z_weixin_service.Res, error) {
	response := new(z_weixin_service.Res)
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}

	parm := new(z_weixin_service.WxParm)
	parm.Type = "image"
	parm.Offset = 0
	parm.Count = 20
	jsonBytes, err := json.Marshal(parm)
	if err != nil {
		return nil, err
	}

	requrl := fmt.Sprintf("%s?access_token=%s", zconfig.SERVICE_APIURL_BATCHGET_MATERIAL, access_token)

	err = zhttp.POSTWithUnmarshal(requrl, jsonBytes, response)
	return response, err
}
