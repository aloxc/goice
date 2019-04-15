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

![ice协议](assert/ice_p.png)

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

| 属性          | 类型     | 长度(字节) | 说明                                                         |
| ------------- | -------- | ---------- | ------------------------------------------------------------ |
| HeaderData    | 结构体   | 14         | 定义响应头                                                   |
| requestId     | int      | 4          | 请求id，每个请求都有不同的id                                 |
| 响应状态码 | byte   | 1       | 响应状态码ReplyStatus |
| 整形后的数据长度 | byte   | 4       | 整形后的数据长度 |
| 编码主版本 | byte   | 1       | 编码版本 |
| 编码次版本 | byte   | 1       | 编码版本 |
| 最终数据长度         | byte或者byte+int | 1或者1+4       | 最终数据长度                                              |
| 最终数据     | string   | 不定       | 数据                                                     |




```
//写完上面的18字节
	// 接下来写 identity.name的长度（见BaseStream的WriteSize方法），接下来写identity.name
	// 接下来写 identity.category的长度（见BaseStream的WriteSize方法），接下来写identity.category（如果为空或者长度为零就不写）
	// Identity的name和category，本里中name=HelloIce,category为空，
	//写完这些数据后buf就有18+1+8+1+0=28字节
	//接下来下facet的，如果facet为空或者长度为0 就写一个为0（byte）的数据到buf后面；
	//如果不为空就封装成facet数组（数组长度为1），然后写数组长度到buf后面，然后循环写每个facet长度和当前facet
	//接下来写调用方法的长度和方法名称，本示例中调用sayHello方法，
	//接下来写一字节的OperationMode，
	//接下来写context的数据（就是指ice.ctx这个context），
	//如果context不为空就写context的size，然后遍历context中key和value，key、value都是字符串，也就是先写key的长度再写key
	//再写value的长度再写value，如此遍历直到遍历完整
	//context为空就读取内置的context（implicitContext，就是我们配置文件内些配置，比如说超时，比如说最大消息长度），
	// 见java代码OutgoingAsync中
	// Ice.ImplicitContextI implicitContext = ref.getInstance().getImplicitContext();
	//            java.util.Map<String, String> prxContext = ref.getContext()
	//如果 implicitContext为空就写 prxContent,否则就写implicitContext和prxContext合并的，但是实际上目前也是空的，。
	//接下来直接写个int 0；
	//

	//连接后发送 18个头，接下来 identity 18 + 1 + 8 + 1
	// + facet 1 = 29
	// + operator(ice_isA) + 1 + 7 = 37
	// + mode + 1 = 38
	// + context = 39
	// + int0 + 4 = 43
	// + encodingVersion + 1 +1 = 45
	// + ::service::HelloService = 45 + 1 + 23 = 69
	//为什么头会改变了呢
	//使用小端
```
即将添加的新功能点
> 添加并发控制
> 添加连接超时
> 添加执行超时
> 添加接受服务端返回的异常
> 支持execute传入json或者Request请求接口
> 支持异步
> 请求及相应统计（增加UI）
> 超时重连或者重新执行
> 加入心跳监测
> 支持在客户端对服务器负载均衡(也允许不配置负载均衡)
> 