package main
//发送40字节的head试试

import (
	"bufio"
	"fmt"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"net"
	"time"
)
var (
	ProtocolMajor         byte = 1
	ProtocolMinor         byte = 0
	ProtocolEncodingMajor byte = 1
	ProtocolEncodingMinor byte = 0

	EncodingMajor byte = 1
	EncodingMinor byte = 1

	//
	// The Ice protocol message types
	//
	RequestMsg            byte = 0

	RequestBatchMsg       byte = 1
	ReplyMsg              byte = 2
	ValidateConnectionMsg byte = 3
	CloseConnectionMsg    byte = 4

	Byte0 byte = 0

)

func main1() {
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
	//	Magic[0],
	//	Magic[1],
	//	Magic[2],
	//	Magic[3],
	//	ProtocolMajor,
	//	ProtocolMinor,
	//	ProtocolEncodingMajor,
	//	ProtocolEncodingMinor,
	//	RequestMsg,
	//	Byte0,                      // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩
	//	Byte0, Byte0, Byte0, Byte0, // Message size (placeholder).
	//	Byte0, Byte0, Byte0, Byte0 }// Request ID (placeholder).

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
	var Magic = []byte{0x49, 0x63, 0x65, 0x50}

	requestHdr1 := []byte{

		Magic[0],
		Magic[1],
		Magic[2],
		Magic[3],
		ProtocolMajor,
		ProtocolMinor,
		ProtocolEncodingMajor,
		ProtocolEncodingMinor,
		RequestMsg,
		1}                    // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩


	var facet string
	var buf =ice.NewIceBuff(rw)
	buf.Write(requestHdr1) // 18字节
	buf.Write(utils.IntToBytes(69))
	buf.Write(utils.IntToBytes(8))
	identity := ice.Identity{
		Name:"HelloIce",
	}
	buf.WriteStr(identity.Name)//18+1+8=27
	buf.WriteStr(identity.Category)//27+1=28

	if len(facet) == 0{
		buf.WriteByte(0)//28+1=29
	}else{
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}
	operator := "ice_isA"
	buf.WriteStr(operator)//29+1+7=37
	var mode byte = 1
	buf.WriteByte(mode)//37+1=38
	context := make(map[string]string)
	buf.WriteStringMap(context)//38+1=39
	//buf.WriteByte(0)//39+1=40
	//buf.WriteByte(0)//40+1=41
	//buf.WriteByte(0)//41+1=42
	//buf.WriteByte(0)//42+1=43
	buf.Write(utils.IntToBytes(30))

	buf.WriteByte(1)//encoding major 43+1=44
	buf.WriteByte(1)//encoding minor 44+1=45


	buf.WriteStr("::service::HelloService")//45+1+23=69

	//数据整形 java 中 BasicStream.endWriteEncaps方法，大约344行，写此后（39位后的）的数据长度，总数据长度减去39
	//还要设置压缩位，见requestHdr里面关于压缩位设置。如果压缩的话，第10位设置为2，并且11 12 13 14设置为压缩后的长度。

	//需要重写requestId，到15 16 17 18 这四位 对应java中 ConnectionI的sendAsyncRequest方法，大约是386行



	buf.Flush()
	time.Sleep(2*time.Second)

	var head = make([]byte,14)//先读取头
	size, err := rw.Read(head)
	fmt.Printf("头 size = %d ,requestId = %d ,msg.size = %d \n" ,size,utils.BytesToInt(head[10:]),utils.BytesToInt(head[10:]))
	var magic = [4]byte{}
	magic[0] = head[0]
	magic[1] = head[1]
	magic[2] = head[2]
	magic[3] = head[3]
	pmj := head[4]
	pmn := head[5]
	emj := head[6]
	emn := head[7]
	rmsg := head[8]
	zip := head[9]
	for i,v:= range magic{
		fmt.Printf("magic[%d]=%d\n",i,v)
	}
	fmt.Printf("协议版本major = %d,minor = %d\n",pmj,pmn)
	fmt.Printf("编码版本major = %d,minor = %d\n",emj,emn)
	fmt.Printf("msg = %d\n",rmsg)
	fmt.Printf("压缩标示 = %d\n",zip)
	fmt.Printf("数据长度 = %d\n",utils.BytesToInt(head[10:]))
	var data = make([]byte,40)//先读取头
	size, err = rw.Reader.Read(data)
	requestId := utils.BytesToInt(data[0:4])
	fmt.Printf("请求ID = %d\n",requestId)


	fmt.Printf("头 size = %d ,requestId = %d ,msg.size = %d \n" ,size,utils.BytesToInt(data[0:4]),utils.BytesToInt(data[0:4]))
	//for size == len(data) {
	//	fmt.Println("size " ,size)
	//}

	//size, err := rw.Read(data)
	//size, err := rw.Read(data)
	//switch {
	//case err == io.EOF:
	//	fmt.Println("读取完成.")
	//case err != nil:
	//	fmt.Println("读取出错")
	//}
	fmt.Println("读取到了 size=",size)
	//fmt.Println("数据长度 len=",len(data))
	//fmt.Println("数据",s)
	time.Sleep(2 * time.Second)
}


