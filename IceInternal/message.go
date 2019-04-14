package IceInternal

import (
	"github.com/aloxc/goice/ice"
	"github.com/siddontang/go/sync2"
)
type MessageType struct {
	//__RequestBatchMsg       byte = 1
	//__ReplyMsg              byte = 2
	//__ValidateConnectionMsg byte = 3
	//__CloseConnectionMsg    byte = 4

}
//https://doc.zeroc.com/ice/3.7/the-ice-protocol/protocol-messages
//Message Header
//Each protocol message has a 14-byte header that is encoded as if it were the following structure:
//请求头数据
type HeadData struct {
	// 目前，协议和编码都是1.0版本。
	Magic             int  //由ASCII编码的值“I”，“c”，“e”，“P”（0x49,0x63,0x65,0x50）组成的四字节幻数
	ProtocolMajor     byte //协议主要版本号
	ProtocolMinor     byte //协议次要版本号
	EncodingMajor     byte //编码主要版本号
	EncodingMinor     byte //编码次要版本号
	MessageType       byte //消息类型 请求0 批量请求1 答复2 验证连接3 关闭连接4
	CompressionStatus byte //消息的压缩状态
	MessageSize       int  //消息的大小（以字节为单位），包括标头
}

//请求正文
type RequestData struct {
	RequestId *sync2.AtomicInt32 //请求id
	Facet     string             //版本控制用
	//一个Ice 对象具有一个特殊的接口，称为它的主接口。此外， Ice 对象还可以提供零个或更多其他接口，称为facets （面）。客户可以在某个对象的各个facets 之间进行挑选，选出它们想要使用的接口。
	//. 每个Ice 对象都有一个唯一的对象标识（object identity）。对象标识是用于把一个对象与其他所有对象区别开来的标识值。Ice 对象模型假定对象标识是全局唯一的，也就是说，在一个Ice 通信域中，不会有两个对
	//象具有相同的对象标识。
	Identity  *ice.Identity   //该对象标识
	Operation string          //操作名称，也就是要调用的方法名称
	Mode      byte            //Ice::OperationMode（0= normal，2= idempotent）的字节表示
	Context   *ice.IceContext //调用上下文
	Param     string          //参数
}

//批量请求正文
type BatchRequestData struct {
	RequestId *sync2.AtomicInt32
}

//相应正文
type ReplyData struct {
	RequestId  int
	StatusCode int
	//Body int
}
