namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

// 微信自定义菜单
struct menu  {
	1:list<button> button
    2:matchrule matchrule
}

// 微信菜单主按钮
struct button  {
	1:string type
	2:string name
	3:string key
	4:string url
	5:string media_id
    6:string appid
    7:string pagepath
	8:list<sub_button> sub_button
}

// 微信菜单子按钮
struct sub_button  {
	1:string type
	2:string name
	3:string key
	4:string url
	5:string media_id
    6:string appid
    7:string pagepath
}

// 微信个性化菜单的选项
struct matchrule  {
	1:string tag_id
	2:i32 sex
	3:string country
	4:string province
	5:string city
	6:string client_platform_type
	7:string language
}