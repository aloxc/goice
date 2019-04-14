package main

//发送40字节的head试试

import (
	"bufio"
	"fmt"
	"github.com/aloxc/goice/IceInternal"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"net"
	"time"
)

var (
	__ProtocolMajor           byte = 1
	__ProtocolMinor           byte = 0
	__Protocol__EncodingMajor byte = 1
	__Protocol__EncodingMinor byte = 0

	__EncodingMajor byte = 1
	__EncodingMinor byte = 1

	//
	// The Ice protocol message types
	//
	__RequestMsg byte = 0

	__RequestBatchMsg       byte = 1
	__ReplyMsg              byte = 2
	__ValidateConnectionMsg byte = 3
	__CloseConnectionMsg    byte = 4

	__Byte0 byte = 0
	__Magic      = [4]byte{0x49, 0x63, 0x65, 0x50}
)

func main() {
	var remoteAddress, _ = net.ResolveTCPAddr("tcp4", "127.0.0.1:1888") //生成一个net.TcpAddr对像。
	var conn, err = net.DialTCP("tcp4", nil, remoteAddress)             //传入协议，本机地址（传了nil），远程地址，获取连接。
	if err != nil { //如果连接失败。则返回。
		fmt.Println("连接出错：", err)
	}
	var remoteIpAddress = conn.RemoteAddr()  //获取IP地址的方法。
	fmt.Println("远程IP地址是：", remoteIpAddress) //输出：220.181.111.188:80
	var localIPAddress = conn.LocalAddr()
	fmt.Println("本地IP地址是：", localIPAddress) //输出：192.168.1.9:45712
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	//requestHdr := []byte{
	//	__Magic[0],
	//	__Magic[1],
	//	__Magic[2],
	//	__Magic[3],
	//	__ProtocolMajor,
	//	__ProtocolMinor,
	//	__Protocol__EncodingMajor,
	//	__Protocol__EncodingMinor,
	//	__RequestMsg,
	//	__Byte0,                      // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩
	//	__Byte0, __Byte0, __Byte0, __Byte0, // Message size (placeholder).
	//	__Byte0, __Byte0, __Byte0, __Byte0 }// Request ID (placeholder).

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

	requestHdr1 := []byte{
		__Magic[0],
		__Magic[1],
		__Magic[2],
		__Magic[3],
		__ProtocolMajor,
		__ProtocolMinor,
		__Protocol__EncodingMajor,
		__Protocol__EncodingMinor,
		__RequestMsg,
		1} // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩
	//还要设置压缩位，见requestHdr里面关于压缩位设置。如果压缩的话，第10位设置为2，并且11 12 13 14设置为压缩后的长度。

	var facet string
	var buf = IceInternal.NewIceBuff(rw)
	var identity ice.Identity
	var context map[string]string
	var mode,pmj,pmn,emj,emn,zip,rmsg byte
	var requestId,size int
	var operator string
	var head ,data []byte

	buf.Write(requestHdr1) // 10字节
	buf.Write(utils.IntToBytes(69)) //size 10 +4 = 14
	buf.Write(utils.IntToBytes(1)) //requestId 14+4=18
	identity = ice.Identity{
		Name: "HelloIce",
	}
	buf.WriteStr(identity.Name)     //18+1+8=27
	buf.WriteStr(identity.Category) //27+1=28

	if len(facet) == 0 {
		buf.WriteByte(0) //28+1=29
	} else {
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}
	operator = "ice_isA"
	buf.WriteStr(operator) //29+1+7=37
	mode = 1
	buf.WriteByte(mode) //37+1=38
	context = make(map[string]string)
	buf.WriteStringMap(context) //38+1=39
	//数据整形 java 中 BasicStream.endWriteEncaps方法，大约344行，写此后（39位后的）的数据长度，总数据长度减去39
	buf.Write(utils.IntToBytes(30))//修正数据长度，是总长度减去写完context后的长度 39 +4 = 43

	buf.WriteByte(1) //encoding major 43+1=44
	buf.WriteByte(1) //encoding minor 44+1=45

	buf.WriteStr("::service::HelloService") //45+1+23=69




	buf.Flush()

	head = make([]byte, 14) //先读取头
	size, err = rw.Read(head)
	var __Magic = [4]byte{}
	__Magic[0] = head[0]
	__Magic[1] = head[1]
	__Magic[2] = head[2]
	__Magic[3] = head[3]
	pmj = head[4]
	pmn = head[5]
	emj = head[6]
	emn = head[7]
	rmsg = head[8]
	zip = head[9]
	fmt.Printf("头 size = %d ,requestId = %d ,msg.size = %d \n", size, utils.BytesToInt(head[10:]), utils.BytesToInt(head[10:]))
	for i, v := range __Magic {
		fmt.Printf("__Magic[%d]=%d\n", i, v)
	}

	fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	fmt.Printf("msg = %d\n", rmsg)
	fmt.Printf("压缩标示 = %d\n", zip)
	fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:]))

	data = make([]byte, 40) //先读取头
	size, err = rw.Read(data)
	requestId = utils.BytesToInt(data[0:4])
	fmt.Printf("请求ID = %d\n", requestId)
	fmt.Println("================")

	time.Sleep(20 * time.Second)

	buf.Write(requestHdr1) // 18字节
	buf.Write(utils.IntToBytes(49))
	buf.Write(utils.IntToBytes(2))

	identity = ice.Identity{
		Name: "HelloIce",
	}
	buf.WriteStr(identity.Name)     //18+1+8=27
	buf.WriteStr(identity.Category) //27+1=28

	if len(facet) == 0 {
		buf.WriteByte(0) //28+1=29
	} else {
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}


	operator = "sayHello"
	buf.WriteStr(operator) //30+1+7=38
	mode = 0
	buf.WriteByte(mode) //38+1=39
	context = make(map[string]string)
	buf.WriteStringMap(context) //39+1=40
	//buf.WriteByte(0)//40+1=41
	//buf.WriteByte(0)//41+1=42
	//buf.WriteByte(0)//42+1=43
	//buf.WriteByte(0)//43+1=44
	//TODO 需要设置实际长度
	buf.Write(utils.IntToBytes(9)) //

	buf.WriteByte(1) //encoding major 44+1=45
	buf.WriteByte(1) //encoding minor 45+1=46

	buf.WriteStr("aa") //46+1+2=49
	buf.Flush()

	head = make([]byte, 25) //先读取头
	size, err = rw.Read(head)
	__Magic = [4]byte{}
	__Magic[0] = head[0]
	__Magic[1] = head[1]
	__Magic[2] = head[2]
	__Magic[3] = head[3]
	pmj = head[4]
	pmn = head[5]
	emj = head[6]
	emn = head[7]
	rmsg = head[8]
	zip = head[9]
	for i, v := range __Magic {
		fmt.Printf("__Magic[%d]=%d\n", i, v)
	}
	fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	fmt.Printf("msg = %d\n", rmsg)
	fmt.Printf("压缩标示 = %d\n", zip)
	fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:14]))
	requestId = utils.BytesToInt(head[14:18])
	fmt.Printf("请求ID = %d\n", requestId)
	var replyStatus uint8 = head[18]
	fmt.Println("响应状态 ", replyStatus)
	var lastSize = utils.BytesToInt(head[19:23])
	fmt.Println("整形后面的数据长度（包括整形4字节） ", lastSize)
	_encodingMajor := head[23]
	_encodingMinor := head[24]
	lastSize = lastSize - 4 - 1 - 1 //4:整形后面包括整形长度，1：主编码版本 ，1：副编码版本
	fmt.Printf("编码版本major = %d,minor = %d\n", _encodingMajor, _encodingMinor)
	fmt.Println("最终数据长度及数据 的长度", lastSize)
	sizeDefine := 0
	realSize := 0
	if lastSize <= 255 {
		//读取一个字节
		sizeDefine = 1
		dataSizeData := make([]byte, sizeDefine)
		size, err = rw.Read(dataSizeData)
		realSize = int(dataSizeData[0])
	} else {
		//读取一个字节-1，跟着4个字节的数据长度（int）
		sizeDefine = 5
		dataSizeData := make([]byte, sizeDefine)
		size, err = rw.Read(dataSizeData)
		realSize = utils.BytesToInt(dataSizeData[1:])
		fmt.Println("长度超过254，读取下这个 -1 是什么%d", dataSizeData[0])
	}
	lastSize = lastSize - sizeDefine
	fmt.Println("计算数据长度是 ", lastSize)
	fmt.Println("真实数据长度是 ", realSize)
	data = make([]byte, lastSize ) //先读取头
	size, err = rw.Read(data)
	fmt.Println("剩余数据长度 = " ,size)
	//fmt.Printf("执行结果[]\n", string(data[:]))
	for in,d := range data{
		fmt.Printf("执行结果[%d][%d]\n",in, d)
	}
	//响应头部10字节 49 63 65 50 01 00 01 00 02 00  // 10
	//数据总长度， 1D 00 00 00   //10+4 = 14
	//请求id   02 00 00 00   //14+4 = 18
	// 一个byte的零 00       //18+1=19   java代码ReplyStatus.replyOK,如果异常就是ReplyStatus.replyUserException
	// 一个int的零 数据长度，最终会整形的（整形是总的数据长度去掉这个位置之前的长度） 0A 00 00 00         //19 + 4 = 23
	// 加上encodingVersion  01 01     //23+2 = 25
	//实际数据
	//数据长度  03
	//数据   62 62 62
	//fmt.Println("数据",s)
	time.Sleep(2 * time.Second)
}
