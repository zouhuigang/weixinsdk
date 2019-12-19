namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service


//返回的素材
struct Item{
	1: string media_id,
	2: string name,
	3:i64 update_time,
	4: string url,
}

struct Res{
	1: string total_count,
	2: string item_count,
	3:list<Item> item,
}

struct WxParm{
	1:string type,
	2:i32 offset,
	3:i32 count,
}

struct MaCount{
	1:i64 voice_count,
	2:i64 video_count,
	3:i64 image_count,
	4:i64 news_count,
}

// type Item struct {
// 	MediaId    string `json:"media_id"`
// 	Name       string `json:"name"`
// 	UpdateTime string `json:"update_time"`
// 	Url        string `json:"url"`
// }

// type Res struct {
// 	TotalCount int64  `json:"total_count"`
// 	ItemCount  int64  `json:"item_count"`
// 	Items      []Item `json:"item"`
// }

// type WxParm struct {
// 	Type   string `json:"type"`
// 	Offset int    `json:"offset"`
// 	Count  int    `json:"count"`
// }

// type MaCount struct {
// 	TotalCount int64 `json:"voice_count"`
// 	VideoCount int64 `json:"video_count"`
// 	ImageCount int64 `json:"image_count"`
// 	NewsCount  int64 `json:"news_count"`
// }