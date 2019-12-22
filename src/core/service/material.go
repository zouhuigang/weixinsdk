package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	zcache "weixinsdk/src/cache"
	zconfig "weixinsdk/src/config"
	z_weixin_service "weixinsdk/src/thrift_file/gen-go/tencent/weixin/service"

	"github.com/zouhuigang/package/zhttp"
)

func (this *WxServiceThrift) MaterialCount() (*z_weixin_service.MaCount, error) {
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
func (this *WxServiceThrift) MaterialList(types string, page int64, pageSize int64) (*z_weixin_service.Res, error) {
	response := new(z_weixin_service.Res)
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 15
	}

	offset := (page - 1) * pageSize
	parm := new(z_weixin_service.WxParm)
	parm.Type = types
	parm.Offset = offset
	parm.Count = pageSize
	jsonBytes, err := json.Marshal(parm)
	if err != nil {
		return nil, err
	}

	requrl := fmt.Sprintf("%s?access_token=%s", zconfig.SERVICE_APIURL_BATCHGET_MATERIAL, access_token)

	err = zhttp.POSTWithUnmarshal(requrl, jsonBytes, response)
	return response, err
}

//上传图片
//curl "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE"
//-F media=@media.file -F  description='{"title":VIDEO_TITLE, "introduction":INTRODUCTION}'
func (this *WxServiceThrift) UpImage(utype string, filename string) (*z_weixin_service.WxImage, error) {
	response := new(z_weixin_service.WxImage)
	// //打开文件
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("media", filepath.Base(filename))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	bodyWriter.Close()

	//upload SERVICE_APIURL_MATERIAL_ADD
	access_token, err := zcache.GetAccessToken()
	if err != nil {
		return nil, err
	}

	requrl := fmt.Sprintf("%s?access_token=%s&type=%s", zconfig.SERVICE_APIURL_MATERIAL_ADD, access_token, utype)
	req, err := http.NewRequest("POST", requrl, bodyBuf)
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())

	// urlQuery := req.URL.Query()
	// if err != nil {
	// 	return nil, err
	// }
	// urlQuery.Add("access_token", access_token)
	// urlQuery.Add("type", utype)
	// req.URL.RawQuery = urlQuery.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	jsonbody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	json.Unmarshal(jsonbody, &response)
	return response, nil
}
