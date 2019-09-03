### 快速开始






### 微信公众平台/微信小程序/服务号/订阅号

>由于切换各种语言写代码时，每次都要重新开发一遍微信这类的接口，感觉非常的麻烦，所以写了这个库，统一下，封装成so或rpc的形式，供其他语言调用。


### 本地测试

	需要将外网的ip添加进白名单（开发->基本配置->ip白名单）


### 微信公众号测试账号申请

[http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login](http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login)


测试url:

	http://594ahg.natappfree.cc/weixin/callback

token:

	weixin


### 搭建微信本地调试环境

国内一家较好的服务商[natapp](https://natapp.cn)。注册后下载对应的客户端，然后在其对应的目录下创建 config.ini 文件：

	[default]
	authtoken=                      # 对应一条隧道的authtoken
	clienttoken=                    # 对应客户端的clienttoken,将会忽略authtoken,若无请留空,
	log=none                        # log 日志文件,可指定本地文件, none=不做记录,stdout=直接屏幕输出 ,默认为none
	loglevel=ERROR                  # 日志等级 DEBUG, INFO, WARNING, ERROR 默认为 DEBUG
	http_proxy=                     # 代理设置 如 http://10.123.10.10:3128 非代理上网用户请务必留空



参考:

[https://blog.csdn.net/qq_34096082/article/details/79985141](https://blog.csdn.net/qq_34096082/article/details/79985141)





### 测试模板消息

模板标题:

	测试模板

模板内容:


	{{content.DATA}}
	{{note.DATA}}
	{{translation.DATA}}


小程序真实的是详细内容：


	充值结果{{keyword1.DATA}}
	电表号{{keyword2.DATA}}
	单号{{keyword3.DATA}}
	金额{{keyword4.DATA}}
	充值金额{{keyword5.DATA}}
	充值时间{{keyword6.DATA}}


服务号模板消息,详细内容:

	{{first.DATA}}
	监控主机：{{keyword1.DATA}}
	警告时间：{{keyword2.DATA}}
	服务名称：{{keyword3.DATA}}
	警告事件：{{keyword4.DATA}}
	警告描述：{{keyword5.DATA}}
	{{remark.DATA}}




### 网页授权

Q:redirect_uri域名与后台配置不一致，错误码:10003

A:

	微信公众平台->网页服务->网页帐号->网页授权获取用户基本信息->修改->将域名粘贴进去(不要加http://)

	dki4r6.natappfree.cc






对应关系:


golang server:

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    //transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

python client:

    transport = TSocket.TSocket('192.168.0.1', 9998)
    //transport = TTransport.TBufferedTransport(transport)
	transport = TTransport.TFramedTransport(transport)
    protocol = TBinaryProtocol.TBinaryProtocol(transport)


php client:

	$socket = new TSocket('127.0.0.1','9998');  
	$socket->setSendTimeout(10000);
	$socket->setRecvTimeout(20000);
	//$transport = new TBufferedTransport($socket);
	$transport = new TFramedTransport($socket); 
	$protocol = new TBinaryProtocol($transport);



### 问题


Q:

	2019/08/07 17:18:26 error processing request: Incorrect frame size (2147549185)


A:

传输方式：这个要和服务器使用的一致，注意

	transportFactory := thrift.NewTBufferedTransportFactory(10000000)

这可能有不同的选项，大部分参考代码中给的都是

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

客户端连接时候一定要与此对应。









### 参阅文档

[http://www.zhongruitech.com/4004855601.html](http://www.zhongruitech.com/4004855601.html)

[https://blog.csdn.net/liuxinmingcode/article/details/45696237](https://blog.csdn.net/liuxinmingcode/article/details/45696237)

[https://studygolang.com/articles/13988](https://studygolang.com/articles/13988)





