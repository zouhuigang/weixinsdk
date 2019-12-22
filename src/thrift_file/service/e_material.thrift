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
	1:i64 total_count,
	2:i64 item_count,
	3:list<Item> item,
}

struct WxParm{
	1:string type,
	2:i64 offset,
	3:i64 count,
}

struct MaCount{
	1:i64 voice_count,
	2:i64 video_count,
	3:i64 image_count,
	4:i64 news_count,
}

struct WxImage{
	1:string type,
	2:string media_id,
	3:i64 created_at,
	4:i32 Errcode
	5:string Errmsg
}