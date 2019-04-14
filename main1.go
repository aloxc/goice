package main

import (
	"bufio"
	"fmt"
	"github.com/aloxc/goice/IceInternal"
	"github.com/aloxc/goice/ice"
	"io"
	"net"
	"strings"
	"time"
)

var (
	protocolMajor         byte = 1
	protocolMinor         byte = 0
	protocolEncodingMajor byte = 1
	protocolEncodingMinor byte = 0

	encodingMajor byte = 1
	encodingMinor byte = 1

	//
	// The Ice protocol message types
	//
	requestMsg            byte = 0

	requestBatchMsg       byte = 1
	replyMsg              byte = 2
	validateConnectionMsg byte = 3
	closeConnectionMsg    byte = 4

	byte0 byte = 0

	magic = [4]byte{0x49, 0x63, 0x65, 0x50}
)

func main() {
	//identity := &ice.Identity{
	//	Name:"HelloIce",
	//}
	//facet := ""
	//mode := 0
	//secure := false
	//encodingVersion := GetDefaultEncodingVersion()
	//protocalVersion := GetDefaultProtocolVersion()
	//connstr :=""
	//if strings.Index(connstr,"default") >= 0 {
	//使用缺省协议，走tcp通道
	//Ice.Default.Host", "127.0.0.1"
	//	tcp -p 1888 -t 60000  60000是超时时间
	//propertyPrefix = null
	//Ice.Default.CollocationOptimized = 1 ，判断最后使用true
	//cacheConnection
	//preferSecure
	//Ice.Default.EndpointSelection = random，EndpointSelectionType也是random
	//Ice.Default.InvocationTimeout调用超时
	//_response = true发送消息
	//basicStream os = null;
	//os.pos(14);
	//os.writeInt(requestId);
	//}else{
	//
	//}
	var remoteAddress, _ = net.ResolveTCPAddr("tcp4", "127.0.0.1:1888") //生成一个net.TcpAddr对像。
	var conn, err = net.DialTCP("tcp4", nil, remoteAddress)             //传入协议，本机地址（传了nil），远程地址，获取连接。
	if err != nil { //如果连接失败。则返回。
		fmt.Println("连接出错：", err)
	}
	var remoteIpAddress = conn.RemoteAddr()  //获取IP地址的方法。
	fmt.Println("远程IP地址是：", remoteIpAddress) //输出：220.181.111.188:80
	var localIPAddress = conn.LocalAddr()
	fmt.Println("本地IP地址是：", localIPAddress) //输出：192.168.1.9:45712
	//content := "IceP"
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	//rw.WriteString("Ice P......? ........HelloIce...sayHello.........zgd7 32553406;id=")
	//rw.WriteByte('I')
	//rw.WriteByte('c')
	//rw.WriteByte('e')
	//rw.WriteByte('P')
	//rw.WriteByte(1)
	//rw.WriteByte(1)
	//rw.WriteByte(1)
	//rw.WriteByte(1)
	//rw.WriteByte(0)
	//rw.WriteByte(0)
	//
	//rw.Write(IntToBytes(10))
	requestHdr := []byte{
		Magic[0],
		Magic[1],
		Magic[2],
		Magic[3],
		protocolMajor,
		protocolMinor,
		protocolEncodingMajor,
		protocolEncodingMinor,
		requestMsg,
		byte0,                      // Compression status.
		byte0, byte0, byte0, byte0, // Message size (placeholder).
		byte0, byte0, byte0, byte0 }// Request ID (placeholder).

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

	var facet string
	var buf =IceInternal.NewIceBuff(rw)
	buf.Write(requestHdr)
	identity := ice.Identity{
		Name:"HelloIce",
	}
	buf.WriteStr(identity.Name)
	buf.WriteStr(identity.Category)

	if len(facet) == 0{
		buf.WriteByte(0)
	}else{
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}
	method := "sayHello"
	buf.WriteStr(method)
	var mode byte = 0
	buf.WriteByte(mode)
	context := make(map[string]string)
	buf.WriteStringMap(context)
	buf.WriteByte(0)
	buf.WriteByte(0)
	buf.WriteByte(0)
	buf.WriteByte(0)

	buf.WriteStr("abcedf")
	buf.WriteString("天啊一和解jijieeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	buf.Flush()
	//rw.Flush()
	for {
		cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			fmt.Println("读取完成.")
		case err != nil:
			fmt.Println("读取出错")
		}

		cmd = strings.Trim(cmd, "\n ")
		fmt.Println(cmd)
		time.Sleep(1000000000)
	}
	time.Sleep(5000000000000000)
}
