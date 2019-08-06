### Cache管理

> cache 维护access_token和jsapi_ticket等一些有有效期的票据通行证。


##### (普通/全局)access_token:

	access_token是公众号的全局唯一接口调用凭据，公众号调用各接口时都需要使用access_token

	注意：是所有接口都需要使用


##### 微信网页授权access_token

	通过网页授权获得的access_token，只能获取到对应的微信用户信息，与微信用户是一对一关系；




