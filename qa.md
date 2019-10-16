Q：
    memcache: cache miss问题，检查发现微信服务器获取access_token时，报错:"errcode":40164,"errmsg":"invalid ip xx.xx.xx.xx, not in whitelist hint: []


A:

    原因:
    
    微信access_token刷新需要添加服务器白名单
    
    解决方法:
    
    登录微信mp后台 -> 开发 / 基本配置 -> 在右侧将上述报出的IP地址添加到"IP白名单"中即可。