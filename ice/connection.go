package ice

import (
	"bufio"
	"fmt"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/utils"
	"github.com/siddontang/go-log/log"
	"net"
	"time"
)

type Connection struct {
	*net.TCPConn
}
type connAndError struct {
	*Connection
	error
}

//连接超时monitor
func connectTimeoutMonitor(network string, address string, timeout int, caech chan *connAndError) {
	log.Info("启动连接监控启动", address)
	<-time.After(time.Duration(timeout) * time.Second)
	caech <- &connAndError{
		error: NewConnectTimeoutError(address, network, timeout),
	}
	log.Info("超时完成")
}
func connect(network string, address string, caech chan *connAndError) {
	//log.Info("准备连接服务器",address)
	var remoteAddress, _ = net.ResolveTCPAddr(network, address)
	var conn, err = net.DialTCP(network, nil, remoteAddress)
	if err != nil {
		caech <- &connAndError{
			error: err,
		}
		return
	}
	c := &Connection{
		TCPConn: conn,
	}
	c.init()
	caech <- &connAndError{
		Connection: c,
	}
	//log.Info("已连接服务器",address)
	return
}

//like this network:tcp,address :127.0.0.1:1888
func Connect(network string, address string, timeout int) (*Connection, error) {
	var caech = make(chan *connAndError)
	go connect(network, address, caech)
	go connectTimeoutMonitor(network, address, timeout, caech)
	var cae *connAndError = <-caech
	return cae.Connection, cae.error
	//var remoteAddress, _ = net.ResolveTCPAddr(network, address)
	//var conn, err = net.DialTCP(network, nil, remoteAddress)
	//if err != nil {
	//	return nil, err
	//}
	//c := &Connection{
	//	TCPConn: conn,
	//}
	//c.init()
	//return c, nil
}
func InitConnection(identity *Identity, name string, conn *net.Conn) {
	//log.Info("连接后要发送一条head命令")
	var facet string
	rw := bufio.NewReadWriter(bufio.NewReader(*conn), bufio.NewWriter(*conn))
	var buf = NewIceBuff(rw)

	total, real := PrepareHead(identity, "", config.ConfigMap[name][config.Module].(string), nil)
	//log.Infof("name[%s]连接池创建好了%d %d %s %s",name,total,real,config.ConfigMap[name][config.Module].(string),identity.name)

	var context map[string]string

	var size int
	var head, data []byte
	var err error
	buf.Write(*GetConnHead())          // 10字节
	buf.Write(utils.IntToBytes(total)) //size 10 +4 = 14
	buf.Write(utils.IntToBytes(1))     //requestId 14+4=18
	buf.WriteStr(identity.name)        //18+1+8=27
	buf.WriteStr(identity.category)    //27+1=28

	if len(facet) == 0 {
		buf.WriteByte(0) //28+1=29
	} else {
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}
	buf.WriteStr(string(config.Ice_isA))         //29+1+7=37
	buf.WriteByte(byte(OperatorModeNonmutating)) //37+1=38
	context = make(map[string]string)
	buf.WriteStringMap(context) //38+1=39
	//数据整形 java 中 BasicStream.endWriteEncaps方法，大约344行，写此后（39位后的）的数据长度，总数据长度减去39
	buf.Write(utils.IntToBytes(real)) //修正数据长度，是总长度减去写完context后的长度 39 +4 = 43

	buf.WriteByte(1) //encoding major 43+1=44
	buf.WriteByte(1) //encoding minor 44+1=45

	//buf.WriteStr("::service::HelloService") //45+1+23=69,只要字节数量一致就可以
	//::module名（slice中定义的）::服务名（slice中定义的interface）
	buf.WriteStr("::" + config.ConfigMap[name][config.Module].(string) + "::" + config.ConfigMap[name][config.Name].(string)) //45+1+23=69
	//log.Infof("string = [%s]", ("::" + config.ConfigMap[name][config.Module].(string) + "::" + config.ConfigMap[name][config.Name].(string)))
	//buf.WriteStr("::goiceinter::Goice") //45+1+23=69
	buf.Flush()
	head = make([]byte, 14) //先读取头
	size, err = rw.Read(head)
	//var __Magic = [4]byte{}
	//__Magic[0] = head[0]
	//__Magic[1] = head[1]
	//__Magic[2] = head[2]
	//__Magic[3] = head[3]
	//pmj := head[4]
	//pmn := head[5]
	//emj := head[6]
	//emn := head[7]
	//rmsg := head[8]
	//zip := head[9]
	//fmt.Printf("头 size = %d ,requestId = %d ,msg.size = %d \n", size, utils.BytesToInt(head[10:]), utils.BytesToInt(head[10:]))
	//fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	//fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	//fmt.Printf("msg = %d\n", rmsg)
	//fmt.Printf("压缩标示 = %d\n", zip)
	//fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:]))

	if err != nil {
		fmt.Println(size)
		panic(err)
	}
	data = make([]byte, 26) //连接的时候
	size, err = rw.Read(data)
	//log.Info("flush reply size : " ,size,utils.BytesToInt32(data[0:4]))
	//fmt.Printf("requestId = %d,数据长度 = %d\n", utils.BytesToInt(data[14:]),utils.BytesToInt(data[10:]))

	//os.Exit(1)
}

func (this *Connection) init() {
	magic := []byte{0x49, 0x63, 0x65, 0x50}
	var msgType byte = 0
	requestHdr := []byte{
		magic[0],
		magic[1],
		magic[2],
		magic[3],
		GetDefaultProtocolVersion().Major,
		GetDefaultProtocolVersion().Minor,
		GetDefaultProtocalEncodingVersion().Major,
		GetDefaultProtocalEncodingVersion().Minor,
		msgType,
		1} // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩
	//还要设置压缩位，见requestHdr里面关于压缩位设置。如果压缩的话，第10位设置为2，并且11 12 13 14设置为压缩后的长度。
	var facet string
	rw := bufio.NewReadWriter(bufio.NewReader(this), bufio.NewWriter(this))
	var buf = NewIceBuff(rw)
	var identity Identity
	var context map[string]string
	var mode byte
	//var pmj, pmn, emj, emn, zip, rmsg byte
	//var requestId int
	var size int
	var operator string
	var head, data []byte
	var err error
	buf.Write(requestHdr)           // 10字节
	buf.Write(utils.IntToBytes(69)) //size 10 +4 = 14
	buf.Write(utils.IntToBytes(1))  //requestId 14+4=18
	identity = Identity{
		name: "HelloIce",
	}
	buf.WriteStr(identity.name)     //18+1+8=27
	buf.WriteStr(identity.category) //27+1=28

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
	buf.Write(utils.IntToBytes(30)) //修正数据长度，是总长度减去写完context后的长度 39 +4 = 43

	buf.WriteByte(1) //encoding major 43+1=44
	buf.WriteByte(1) //encoding minor 44+1=45

	buf.WriteStr("::service::HelloService") //45+1+23=69
	buf.Flush()

	head = make([]byte, 14) //先读取头
	size, err = rw.Read(head)
	if err != nil {
		fmt.Println(size)
		panic(err)
	}
	var __Magic = [4]byte{}
	__Magic[0] = head[0]
	__Magic[1] = head[1]
	__Magic[2] = head[2]
	__Magic[3] = head[3]
	//pmj = head[4]
	//pmn = head[5]
	//emj = head[6]
	//emn = head[7]
	//rmsg = head[8]
	//zip = head[9]
	//fmt.Printf("头 size = %d ,requestId = %d ,msg.size = %d \n", size, utils.BytesToInt(head[10:]), utils.BytesToInt(head[10:]))
	//for i, v := range __Magic {
	//	fmt.Printf("__Magic[%d]=%d\n", i, v)
	//}
	//
	//fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	//fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	//fmt.Printf("msg = %d\n", rmsg)
	//fmt.Printf("压缩标示 = %d\n", zip)
	//fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:]))

	data = make([]byte, 40) //先读取头
	size, err = rw.Read(data)
	//requestId = utils.BytesToInt(data[0:4])
	//fmt.Printf("请求ID = %d\n", requestId)
	//fmt.Println("================")
}
