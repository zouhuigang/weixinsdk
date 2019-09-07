namespace go tencent.weixin.service
namespace php tencent.weixin.service
namespace py tencent.weixin.service
/**
map(t,t): 键类型为t，值类型为t的kv对，键不容许重复。对应c++中的map, Java的HashMap, PHP 对应 array, Python/Ruby 的dictionary
 http://www.cpper.cn/2016/03/18/develop/Thrift-The-Missing-Guide/
 http://ju.outofmemory.cn/entry/263563
bool 布尔型
byte ８位整数
i16  16位整数
i32  32位整数
i64  64位整数
double 双精度浮点数
string 字符串
binary 字节数组
list<i16> List集合，必须指明泛型
map<string, string> Map类型，必须指明泛型
set<i32> Set集合，必须指明泛型 
 */
struct Article{
 1: i32 id, 
 2: string title,
 3: string content,
 4: string author,
}