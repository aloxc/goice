package ice

import (
	"bufio"
	"fmt"
	"github.com/aloxc/goice/utils"
	"net"
)

type Connection struct {
	*net.TCPConn
}

//like this network:tcp,address :127.0.0.1:1888
func Connect(network string, address string) (*Connection, error) {
	var remoteAddress, _ = net.ResolveTCPAddr(network, address)
	var conn, err = net.DialTCP(network, nil, remoteAddress)
	if err != nil {
		return nil, err
	}
	c := &Connection{
		TCPConn: conn,
	}
	c.init()
	return c, nil
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
	buf.Write(requestHdr)          // 10字节
	buf.Write(utils.IntToBytes(69)) //size 10 +4 = 14
	buf.Write(utils.IntToBytes(1))  //requestId 14+4=18
	identity = Identity{
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
	buf.Write(utils.IntToBytes(30)) //修正数据长度，是总长度减去写完context后的长度 39 +4 = 43

	buf.WriteByte(1) //encoding major 43+1=44
	buf.WriteByte(1) //encoding minor 44+1=45

	buf.WriteStr("::service::HelloService") //45+1+23=69
	buf.Flush()



	head = make([]byte, 14) //先读取头
	size, err = rw.Read(head)
	if err != nil{
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