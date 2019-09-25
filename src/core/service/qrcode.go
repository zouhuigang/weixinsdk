package service

import (
	"fmt"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	"weixinsdk/src/logger"

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

func (this *WxServiceThrift) QrcodeShow(qrJsonBytes []byte) (string, error) {
	//menuJsonByte := bytes.NewReader(menuJsonBytes)
	logger.MyLogger.Debug("QrcodeShow:" + string(qrJsonBytes))

	//获取ticket
	qrResponse, err := qrcodeCreate(qrJsonBytes)
	if err != nil {
		return "", err
	}

	qrUrl := fmt.Sprintf("%s?ticket=%s", zconfig.SERVICE_APIURL_QRCODE_SHOW, qrResponse.Ticket)

	return qrUrl, nil
}
