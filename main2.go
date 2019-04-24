package main

//发送40字节的head试试

func main2() {
	//var conn, err = ice.Connect("tcp4", "127.0.0.1:1888", 0)
	//if err != nil { //如果连接失败。则返回。
	//	fmt.Println("连接出错：", err)
	//}
	//rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	//var facet string
	//var buf = ice.NewIceBuff(rw)
	//var context map[string]string
	//var mode, pmj, pmn, emj, emn, zip, rmsg byte
	//var requestId, size int
	//var head, data []byte
	//
	////time.Sleep(20 * time.Second)
	//
	//buf.WriteHead()
	//buf.WriteTotalSize(49)
	//buf.WriteRequestId(3)
	//buf.WriteIdentity(ice.GetIdentity("HelloIce", ""))
	//buf.WriteFacet(facet)
	//buf.WriteOperator("sayHello")
	//mode = 0
	//buf.WriteMode(mode) //38+1=39
	//context = make(map[string]string)
	//buf.WriteContext(context) //39+1=40
	////需要设置实际长度,这个值是总数据长度减去设置完context后的长度
	//buf.WriteRealSize(9)
	//buf.WriteEncodingVersion(ice.GetDefaultEncodingVersion())
	//buf.WriteStr("aa") //46+1+2=49
	//buf.Flush()
	//
	//head = make([]byte, 25) //先读取头
	//size, err = rw.Read(head)
	//var __Magic = [4]byte{}
	//__Magic[0] = head[0]
	//__Magic[1] = head[1]
	//__Magic[2] = head[2]
	//__Magic[3] = head[3]
	//pmj = head[4]
	//pmn = head[5]
	//emj = head[6]
	//emn = head[7]
	//rmsg = head[8]
	//zip = head[9]
	//for i, v := range __Magic {
	//	fmt.Printf("__Magic[%d]=%d\n", i, v)
	//}
	//fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	//fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	//fmt.Printf("msg = %d\n", rmsg)
	//fmt.Printf("压缩标示 = %d\n", zip)
	//fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:14]))
	//requestId = utils.BytesToInt(head[14:18])
	//fmt.Printf("请求ID = %d\n", requestId)
	//var replyStatus uint8 = head[18]
	//fmt.Println("响应状态 ", replyStatus)
	//var lastSize = utils.BytesToInt(head[19:23])
	//fmt.Println("整形后面的数据长度（包括整形4字节） ", lastSize)
	//_encodingMajor := head[23]
	//_encodingMinor := head[24]
	//lastSize = lastSize - 4 - 1 - 1 //4:整形后面包括整形长度，1：主编码版本 ，1：副编码版本
	//fmt.Printf("编码版本major = %d,minor = %d\n", _encodingMajor, _encodingMinor)
	//fmt.Println("最终数据长度及数据 的长度", lastSize)
	//sizeDefine := 0
	//realSize := 0
	//if lastSize <= 255 {
	//	//读取一个字节
	//	sizeDefine = 1
	//	dataSizeData := make([]byte, sizeDefine)
	//	size, err = rw.Read(dataSizeData)
	//	realSize = int(dataSizeData[0])
	//} else {
	//	//读取一个字节-1，跟着4个字节的数据长度（int）
	//	sizeDefine = 5
	//	dataSizeData := make([]byte, sizeDefine)
	//	size, err = rw.Read(dataSizeData)
	//	realSize = utils.BytesToInt(dataSizeData[1:])
	//	fmt.Println("长度超过254，读取下这个 -1 是什么 %V", dataSizeData[0])
	//}
	//lastSize = lastSize - sizeDefine
	//fmt.Println("计算数据长度是 ", lastSize)
	//fmt.Println("真实数据长度是 ", realSize)
	//data = make([]byte, lastSize) //先读取头
	//size, err = rw.Read(data)
	//fmt.Println("剩余数据长度 = ", size)
	////fmt.Printf("执行结果[]\n", string(data[:]))
	//for in, d := range data {
	//	fmt.Printf("执行结果[%d][%d]\n", in, d)
	//}
	////响应头部10字节 49 63 65 50 01 00 01 00 02 00  // 10
	////数据总长度， 1D 00 00 00   //10+4 = 14
	////请求id   02 00 00 00   //14+4 = 18
	//// 一个byte的零 00       //18+1=19   java代码ReplyStatus.replyOK,如果异常就是ReplyStatus.replyUserException
	//// 一个int的零 数据长度，最终会整形的（整形是总的数据长度去掉这个位置之前的长度） 0A 00 00 00         //19 + 4 = 23
	//// 加上encodingVersion  01 01     //23+2 = 25
	////实际数据
	////数据长度  03
	////数据   62 62 62
	////fmt.Println("数据",s)
	//time.Sleep(2 * time.Second)
}
