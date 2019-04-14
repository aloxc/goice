# 1、ice请求协议
## 1.1、ice连接发需要发送一条消息，其实是把需要调用的ice服务类传到服务器上
协议格式如下：
| 属性          | 类型     | 长度(字节) | 说明                                                         |
| ------------- | -------- | ---------- | ------------------------------------------------------------ |
| HeaderData    | 结构体   | 14         | 定义请求头                                                   |
| requestId     | int      | 4          | 请求id，每个请求都有不同的id                                 |
| Ice::Identity | 结构体   | 不定       | 对象标识，该表示是指启动ice的时候定义好的，例如启动ice的时候定义为Identity id = Util.stringToIdentity("HelloIce"); |
| facet         | string[] | 不定       | ICE版本控制标识                                              |
| operation     | string   | 不定       | 方法名称                                                     |
| mode          | byte     | 1          | 模式，枚举值 OperationMode 这个枚举类中的                    |
| context       | map      | 不定       | 上下文，(string,string)hashmap                               |
| params        |          | 不定       | 请求参数                                                     |

```golang
struct HeaderData
{
    int  magic;
    byte protocolMajor;
    byte protocolMinor;
    byte encodingMajor;
    byte encodingMinor;
    byte messageType;
    byte compressionStatus;
    int  messageSize;
}
```
    
# 2、ice响应协议