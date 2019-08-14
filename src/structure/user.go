package structure

/*
{
    "subscribe": 1,
    "openid": "o5im3vwMdiOSgy9MGxc3s5ZD-QIM",
    "nickname": "whhdh",
    "sex": 1,
    "language": "zh_CN",
    "city": "闵行",
    "province": "上海",
    "country": "中国",
    "headimgurl": "http://thirdwx.qlogo.cn/mmopen/OszSfzGiaEYNLlxOm80ibf9JCk5LZtbULVVRtPb4Tl43wHRQSP6sMDzeLicBcv4qhVZibHu5HlChLQCXibrJsOn3HNjCTkRdU0U4F/132",
    "subscribe_time": 1565443434,
    "remark": "",
    "groupid": 0,
    "tagid_list": [],
    "subscribe_scene": "ADD_SCENE_QR_CODE",
    "qr_scene": 0,
    "qr_scene_str": ""
}
*/
type UserInfo struct {
	Subscribe int    `json:"subscribe"`
	OpenId    string `json:"openid"`   // 用户的唯一标识
	Nickname  string `json:"nickname"` // 用户昵称
	Sex       int    `json:"sex"`      // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City      string `json:"city"`     // 普通用户个人资料填写的城市
	Language  string `json:"language"`
	Province  string `json:"province"` // 用户个人资料填写的省份
	Country   string `json:"country"`  // 国家，如中国为CN

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空
	Headimgurl string `json:"headimgurl,omitempty"`

	Subscribe_time  int64    `json:"subscribe_time"`
	Remark          string   `json:"remark"`
	Groupid         int64    `json:"groupid"`
	Tagid_list      []string `json:"tagid_list"`
	Subscribe_scene string   `json:"subscribe_scene"`
	Qr_scene        int      `json:"qr_scene"`
	Qr_scene_str    string   `json:"qr_scene_str"`

	// 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Privilege []string `json:"privilege"`

	// 用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的unionid是唯一的。
	UnionId string `json:"unionid"`
}
