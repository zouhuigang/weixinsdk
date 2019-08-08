#!/usr/bin/env python
# -*- coding: utf-8 -*-

'''
pip install thrift
'''
import sys
sys.path.append('./gen-py')

from weixin import WxServiceThrift

from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol


def run():
    try:
        transport = TSocket.TSocket('127.0.0.1', 3333)
        transport = TTransport.TFramedTransport(transport)
        protocol = TBinaryProtocol.TBinaryProtocol(transport)

        client = WxServiceThrift.Client(protocol)
        transport.open()
        
        s = WxServiceThrift.Article()
        s.id = 1
        s.title = "插入一篇测试文章"
        s.content = "我就是这篇文章内容"
        s.author = "zouhuigang"
        client.put(s)
        print("success")
        transport.close()
    except Thrift.TException as ex:
        print("%s" % (ex.message))


if __name__ == '__main__':
    run()
  