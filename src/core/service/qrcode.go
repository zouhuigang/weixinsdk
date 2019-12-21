package service

import (
	"fmt"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service"

	"github.com/zouhuigang/package/zhttp"
)

/*
获取带参数的二维码的过程包括两步，
1.首先创建二维码ticket，
2.然后凭借ticket到指定URL换取二维码
*/

/*{"ticket":"gQH47joAAAAAAAAAASxodHRwOi8vd2VpeGluLnFxLmNvbS9xL2taZ2Z3TVRtNzJXV1Brb3ZhYmJJAAIEZ23sUwMEmm3sUw==",
"expire_seconds":60,
"url":"http://weixin.qq.com/q/kZgfwMTm72WWPkovabbI"}
*/

type QrResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
}

func qrcodeCreate(qrJsonBytes []byte) (*QrResponse, error) {
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}

	requrl := fmt.Sprintf("%s?access_token=%s", zconfig.SERVICE_APIURL_QRCODE_CREATE, access_token)
	response := new(QrResponse)
	err = zhttp.POSTWithUnmarshal(requrl, qrJsonBytes, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (this *WxServiceThrift) QrcodeShow(qrJsonBytes []byte) (*z_weixin_service.QrRespone, error) {
	//menuJsonByte := bytes.NewReader(menuJsonBytes)
	response := z_weixin_service.NewQrRespone()
	response.ReqContent = string(qrJsonBytes) //请求的内容
	logger.MyLogger.Debug("QrcodeShow:" + string(qrJsonBytes))

	//获取ticket
	qrResponse, err := qrcodeCreate(qrJsonBytes)
	if err != nil {
		return response, err
	}

	qrUrl := fmt.Sprintf("%s?ticket=%s", zconfig.SERVICE_APIURL_QRCODE_SHOW, qrResponse.Ticket)
	response.URL = qrUrl                                     //二维码网址
	response.Content = qrResponse.Url                        //二维码内容
	response.ExpireSeconds = int32(qrResponse.ExpireSeconds) //该二维码有效时间，以秒为单位。 最大不超过2592000(即30天)

	logger.MyLogger.Debug("QrcodeShow content:" + string(response.Content))
	return response, nil
}
