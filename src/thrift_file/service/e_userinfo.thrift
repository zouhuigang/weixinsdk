namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service

//用户信息
struct UserInfo{
	1:string openid
	2:string nickname
	3:i32 sex
	4:string province
	5:string city
	6:string country
	7:string headimgurl
	8:list<string> privilege
	9:string unionid
}