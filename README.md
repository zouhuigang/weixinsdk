### 微信公众平台/微信小程序/服务号/订阅号


### 本地测试

	需要将外网的ip添加进白名单（开发->基本配置->ip白名单）


### 微信公众号测试账号申请

[http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login](http://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login)


由于切换各种语言写代码时，每次都要重新开发一遍微信这类的接口，感觉非常的麻烦，所以写了这个库，统一下，封装成so或rpc的形式，供其他语言调用。


    go get -u github.com/FiloSottile/gvt



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





