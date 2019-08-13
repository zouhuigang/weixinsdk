### thrift_file 用于存放 thrift 的 IDL 文件： *.thrift

>此处的代码除了.thrift外，其他都是根据.thrift文件生成的。

生成go代码:

	thrift -out .. --gen go example.thrift


进入 thrift_file 目录执行，就会在 thrift_file 的同级目录下生成 golang 的包：example，其中 format_data-remote 是生成的测试代码可以不用特别关注

---


	thrift -r --gen go weixin.thrift
	thrift -r --gen php weixin.thrift
	thrift -r --gen py weixin.thrift
	thrift -r --gen php:server weixin.thrift #生成PHP服务端接口代码有所不一样



https://www.cnblogs.com/exceptioneye/p/4985598.html


### 问题汇总


Q:
	golang代码报错：not enough arguments in call to oprot.Flush


A：

	thrift版本与 go get git.apache.org/thrift.git/lib/go/thrift 下载下来的库版本是不同导致的。

	thrift --version:
	Thrift version 0.10.0

	cd D:\workspacego\src\git.apache.org\thrift.git
	git tag
	0.10.0
	0.11.0
	0.2.0
	0.3.0
	0.4.0
	0.5.0
	0.6.0
	0.6.1
	0.7.0
	0.8.0
	0.9.0
	0.9.1
	0.9.2
	0.9.3
	0.9.3.1
	thrift-0.2.0
	thrift-0.3.0
	thrift-0.4.0
	thrift-0.5.0
	thrift-0.6.0
	thrift-0.6.1
	thrift-0.7.0
	thrift-0.8.0
	thrift-0.9.0
	v0.12.0

	检出:
	git checkout 0.10.0

	再次运行:
	go run main.go
	即可解决问题。



### 参考文档

[https://blog.csdn.net/lanyang123456/article/details/80372977](https://blog.csdn.net/lanyang123456/article/details/80372977)
	
	

