package structure

import "encoding/xml"

const (
	WeMsgTypeText       = "text"
	WeMsgTypeImage      = "image"
	WeMsgTypeVoice      = "voice"
	WeMsgTypeVideo      = "video"
	WeMsgTypeShortVideo = "shortvideo"
	WeMsgTypeLocation   = "location"
	WeMsgTypeLink       = "link"
	WeMsgTypeEvent      = "event"

	WeEventSubscribe          = "subscribe"
	WeEventUnsubscribe        = "unsubscribe"
	WeTransferCustomerService = "transfer_customer_service"
)

//https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140453
//https://github.com/skynology/wechat/blob/3b285c664f877bc4d1db433f949ada4b2afd3d0c/corp/request.go
//https://github.com/bmob/goLib/blob/0447b3e9cf79b92f72e7c355d23f2c29d07555be/chanxuehong/wechat/corp/message/request/msg.go
//https://segmentfault.com/a/1190000019752111?utm_source=tag-newest
//微信服务器推送过来的消息(事件)的通用消息头，不放MsgId的目的是，回复的时候，不需要带msgid
type CommonMessageHeader struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`
	FromUserName string `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`
	MsgType      string `xml:"MsgType"      json:"MsgType"`
	ArticleCount int32  `xml:"ArticleCount"   json:"ArticleCount"`
}

// //文本消息,XMLNAME是最外面的名称
type WxMsgTxt struct {
	XMLName xml.Name `xml:"xml"`
	CommonMessageHeader
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

//客服消息
type WxMsgTransferCustomerService struct {
	XMLName xml.Name `xml:"xml"`
	CommonMessageHeader
	KfAccount string `xml:"kfAccount" json:"kfAccount"` // 指定客服
}

type Image struct {
	MediaId string `xml:"MediaId" json:"MediaId"`
}
type Articles struct {
	XMLName xml.Name `xml:"Articles"`
	Item    []Item
}
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"Title" json:"Title"`
	Description string   `xml:"Description" json:"Description"`
	PicUrl      string   `xml:"PicUrl" json:"PicUrl"`
	Url         string   `xml:"Url" json:"Url"`
}

type WxAutoReply struct {
	XMLName xml.Name `xml:"xml"`
	CommonMessageHeader
	Content  string   `xml:"Content" json:"Content"` // 文本消息内容
	Image    Image    `xml:"Image" json:"Image"`
	Articles Articles `xml:"Articles" json:"Articles"`
}

// //语音消息
// type WxMsgVoice struct {
// 	CommonMessageHeader
// 	MediaId string `xml:"MediaId" json:"MediaId"` // 语音媒体文件id，可以调用获取媒体文件接口拉取数据
// 	Format  string `xml:"Format"  json:"Format"`  // 语音格式，如amr，speex等
// }

// //视频消息 XMLName struct{} `xml:"xml" json:"-"`
// type WxMsgVideo struct {
// 	CommonMessageHeader
// 	MediaId      string `xml:"MediaId"      json:"MediaId"`      // 视频媒体文件id，可以调用获取媒体文件接口拉取数据
// 	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据
// }

// //小视频消息
// type WxMsgShortVideo struct {
// 	CommonMessageHeader
// 	MediaId      string `xml:"MediaId"      json:"MediaId"`      //视频消息媒体id，可以调用获取临时素材接口拉取数据。
// 	ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` //视频消息缩略图的媒体id，可以调用获取临时素材接口拉取数据。
// }

// //地理位置消息
// type WxMsgLocation struct {
// 	XMLName struct{} `xml:"xml" json:"-"`
// 	CommonMessageHeader

// 	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
// 	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
// 	Scale     int     `xml:"Scale"      json:"Scale"`      // 地图缩放大小
// 	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
// }

// //链接消息
// type WxMsgLink struct {
// 	CommonMessageHeader
// 	Title       string `xml:"Title" json:"Title"`             //消息标题
// 	Description string `xml:"Description" json:"Description"` //消息描述
// 	Url         string `xml:"Url" json:"Url"`                 //消息链接
// }

// // 微信服务器推送过来的消息(事件)的合集.
// type MixedMessage struct {
// 	//通用
// 	CommonMessageHeader

// MsgId        int64  `xml:"MsgId"   json:"MsgId"`        // 消息id，64位整型
// AgentId      int64  `xml:"AgentID"      json:"AgentID"` //跳转链接时所在的企业应用ID
// 	//文本消息
// 	Content string `xml:"Content" json:"Content"` // 文本消息内容
// 	// 图片消息
// 	PicURL  string `xml:"PicUrl"  json:"PicUrl"`
// 	MediaId string

// 	// 音频消息
// 	Format string

// 	// 视频或短视频
// 	ThumbMediaId string

// 	// 地理位置消息
// 	Location_X float64
// 	Location_Y float64
// 	Scale      int
// 	Label      string

// 	// 链接消息
// 	Title       string
// 	Description string
// 	Url         string

// 	// 事件
// 	Event string
// }

// type CData struct {
// 	Val string `xml:",cdata"`
// }

// type WechatReply struct {
// 	XMLName      xml.Name `xml:"xml"`
// 	ToUserName   *CData
// 	FromUserName *CData
// 	CreateTime   int64
// 	MsgType      *CData
// 	Content      *CData `xml:",omitempty"`
// }
