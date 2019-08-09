### 微信服务号官方的接口


	https://mp.weixin.qq.com/advanced/advanced?action=table&token=1926782267&lang=zh_CN


##### 获取access_token：

	access_token=client.GetAccessToken()

返回:

	
	24_TiW6PEs1AwzCLIYsQeQI-SaO9dEkinO-YkHFQ_mlhT4vuOCY9cPq02s8Tbl5Bb0b_6UAcsLzXa-JWDhEmneYfOAfegtNSYFhG5QvT5EmlAlKJkCueT4ma8h_ypJ-HQImlKCzCfsNVzbffod_TTVeAHADDF


##### 获取jsapi_ticket

	ticket=client.GetJsapiTicket()

返回:

	HoagFKDcsGMVCIY2vOjf9t8uL4QYTDQSyx-oq2VBxxW_WqGegW5TzcQtFRnLdJElGJPUBSfZMZqy6Zh3XV31uw


##### 微信扫一扫功能

	JsapiSign("http://c3.ab.51tywy.com/qrcode/test/test.html")

返回:

	JsapiSignData(
		jsapi_ticket='HoagFKDcsGMVCIY2vOjf9t8uL4QYTDQSyx-oq2VBxxW_WqGegW5TzcQtFRnLdJElGJPUBSfZMZqy6Zh3XV31uw', noncestr='pmpqqmim2y362ccmrsrwmk3zbzrhwmt8', 
		timestamp=1565323172, 
		url='http://c3.ab.51tywy.com/qrcode/test/test.html', 
		sign='5da1e238185fdfd5925d1a731c789a8f0f7da93f', 
		appid='wx5f00c646abe6af91'
	)