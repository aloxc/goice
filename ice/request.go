package ice

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/aloxc/goice/utils"
	"io"
	"sync/atomic"
	"time"
)

type Request struct {
	Method string
	Params map[string]string
}

type reqeustErrorAndData struct {
	err  error
	data []byte
}

//构造用于execute方法的请求
func NewReqeust(method string, params map[string]string) *Request {
	return &Request{
		Method: method,
		Params: params,
	}
}
func (this *Request) String() string {
	r := make(map[string]interface{})
	r["method"] = this.Method
	r["params"] = this.Params
	if bytes, err := json.Marshal(r); err == nil {
		return string(bytes)
	}
	return ""
}

var requestId int32 = 2

type IceRequest struct {
	head      *[]byte // 10 字节 0 + 10 = 10
	totalSize int     //所有数据长度 4 字节 10 + 4 = 14
	requestId int     //请求id 4 字节 14 + 4 = 18
	*Identity         //该对象标识  1 + xx + 1 + 0 = 10 字节  18 + 2 + xx = 20 + xx
	//一个Ice 对象具有一个特殊的接口，称为它的主接口。此外， Ice 对象还可以提供零个或更多其他接口，称为facets （面）。客户可以在某个对象的各个facets 之间进行挑选，选出它们想要使用的接口。
	//. 每个Ice 对象都有一个唯一的对象标识（object identity）。对象标识是用于把一个对象与其他所有对象区别开来的标识值。Ice 对象模型假定对象标识是全局唯一的，也就是说，在一个Ice 通信域中，不会有两个对
	//象具有相同的对象标识。
	Facet           string            //版本控制用 1 字节 1 + 20 + xx = 21 + xx
	Operator        string            //操作名称，也就是要调用的方法名称 yy 字节 21 + xx + yy
	OperatorMode                      // 1 字节 1 + 21 + xx + yy = 22 + xx + yy
	realSize        int               // 4 字节 4 + 22 + xx + yy = 26 + xx + yy
	Context         map[string]string //调用上下文 zz 字节
	encodingVersion *EncodingVersion  // 2 字节
	Params          interface{}
}

func NewIceRequest(identity *Identity, mode OperatorMode, operator string, context map[string]string, params interface{}) *IceRequest {
	return &IceRequest{
		head:            GetHead(),
		Operator:        operator,
		OperatorMode:    mode,
		encodingVersion: GetDefaultEncodingVersion(),
		Params:          params,
		Identity:        identity,
		Context:         context,
	}
}

//准备把所有设置都放到这个方法中，先Prepare下，然后再调用组装数据的，最后就是执行this.Flush
func (this *IceRequest) DoRequest(responseType ResponseType) ([]byte, error) {
	var timeout int = 5
	atomic.AddInt32(&requestId, 1)
	this.requestId = int(requestId)
	var conn, err = Connect("tcp4", "127.0.0.1:1888")
	if err != nil { //如果连接失败。则返回。
		fmt.Println("连接出错：")
		return nil, err
	}
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	var buf = NewIceBuff(rw)
	//time.Sleep(20 * time.Second)
	total, real := buf.Prepare(this.Identity, this.Facet, this.Operator, this.Params, this.Context)
	buf.Write(*this.head)
	buf.WriteTotalSize(total)
	buf.WriteRequestId(this.requestId)
	buf.WriteIdentity(this.Identity)
	buf.WriteFacet(this.Facet)
	buf.WriteOperator(this.Operator)
	buf.WriteByte(byte(this.OperatorMode))
	buf.WriteContext(this.Context)
	buf.WriteRealSize(real)
	buf.WriteEncodingVersion(this.encodingVersion)
	switch this.Params.(type) {
	case string:
		buf.WriteStr(this.Params.(string))
	case bool:
		buf.Write(utils.BoolToBytes(this.Params.(bool)))
	case int8:
		buf.Write(utils.Int8ToBytes(this.Params.(int8)))
	case int16:
		buf.Write(utils.Int16ToBytes(this.Params.(int16)))
	case int:
		buf.Write(utils.IntToBytes(this.Params.(int)))
	case int32:
		buf.Write(utils.Int32ToBytes(this.Params.(int32)))
	case int64:
		buf.Write(utils.Int64ToBytes(this.Params.(int64)))
	case float32:
		buf.Write(utils.Float32ToByte(this.Params.(float32)))
	case float64:
		buf.Write(utils.Float64ToByte(this.Params.(float64)))
	case *Request:
		request := this.Params.(*Request)
		buf.WriteStr(request.Method)
		buf.WriteStringMap(request.Params)
	}
	buf.Flush()
	fmt.Println("请求已经发送")
	//var timeoutCh chan int

	errAndData := make(chan *reqeustErrorAndData)
	go request(conn.RemoteAddr().String(), rw, responseType, this.Params, errAndData)
	go timeoutMonitor(conn.RemoteAddr().String(), "", timeout, this.Params, errAndData)
	ed := <-errAndData
	return ed.data, ed.err

}
func request(address string, rw io.ReadWriter, responseType ResponseType, params interface{}, errAndData chan *reqeustErrorAndData) {
	var size, lastSize int
	var head, data []byte
	head = make([]byte, 25) //先读取头
	size, err := rw.Read(head)
	if err != nil {
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
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
	//for i, v := range __Magic {
	//	fmt.Printf("__Magic[%d]=%d\n", i, v)
	//}
	//fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	//fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	//fmt.Printf("msg = %d\n", rmsg)
	//fmt.Printf("压缩标示 = %d\n", zip)
	//fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:14]))
	if utils.BytesToInt(head[10:14]) == size || responseType == ResponseType_Void { //void 的方法，没有返回,也就只会返回25个字节的数据
		errAndData <- &reqeustErrorAndData{
			err:  nil,
			data: nil,
		}
		return
	}
	//requestId = utils.BytesToInt(head[14:18])
	//fmt.Printf("请求ID = %d\n", requestId)
	var replyStatus uint8 = head[18]
	fmt.Println("响应状态 ", replyStatus)
	switch replyStatus {
	case 1:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewUserError(address, "", string(data), params),
			data: nil,
		}
	case 2:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewObjectNotExistsError(address, "", string(data), params),
			data: nil,
		}
	case 3:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewFacetNotExistsError(address, "", string(data), params),
			data: nil,
		}
		return
	case 4:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewOperatorNotExistsError(address, "", string(data), params),
			data: nil,
		}
		return
	case 5:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewIceServerError(address, "", string(data), params),
			data: nil,
		}
	case 6:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewUserError(address, "", string(data), params),
			data: nil,
		}
	case 7: //用户异常
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			fmt.Println("读取异常信息异常")
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		userUnknownError := NewUserUnknownError(address, "", string(data), params)
		errAndData <- &reqeustErrorAndData{
			err:  userUnknownError,
			data: nil,
		}
		return

	}
	lastSize = utils.BytesToInt(head[19:23])
	//fmt.Println("整形后面的数据长度（包括整形4字节） ", lastSize)
	//_encodingMajor := head[23]
	//_encodingMinor := head[24]
	lastSize = lastSize - 4 - 1 - 1 //4:整形后面包括整形长度，1：主编码版本 ，1：副编码版本
	//fmt.Printf("编码版本major = %d,minor = %d\n", _encodingMajor, _encodingMinor)
	//fmt.Println("最终数据长度及数据 的长度", lastSize)

	if responseType == ResponseType_Bool {
		data := make([]byte, 1)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	} else if responseType == ResponseType_Int8 {
		data := make([]byte, 1)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	} else if responseType == ResponseType_Int16 {
		data := make([]byte, 2)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	} else if responseType == ResponseType_Int {
		data := make([]byte, 4)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	} else if responseType == ResponseType_Int64 {
		data := make([]byte, 8)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	} else if responseType == ResponseType_Float32 {
		data := make([]byte, 4)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	} else if responseType == ResponseType_Float64 {
		data := make([]byte, 8)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data,
		}
		return
	}
	sizeDefine := 0
	//realSize := 0
	//字符串的
	if lastSize <= 255 {
		//读取一个字节
		sizeDefine = 1
		dataSizeData := make([]byte, sizeDefine)
		size, err = rw.Read(dataSizeData)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		//realSize = int(dataSizeData[0])
	} else {
		//读取一个字节-1，跟着4个字节的数据长度（int）
		sizeDefine = 5
		dataSizeData := make([]byte, sizeDefine)
		size, err = rw.Read(dataSizeData)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		//realSize = utils.BytesToInt(dataSizeData[1:])
		//fmt.Println("长度超过254，读取下这个 -1 是什么%d", dataSizeData[0])
	}
	//fmt.Println("lastSize=",lastSize)
	lastSize = lastSize - sizeDefine
	//fmt.Println("sizeDefine=" ,sizeDefine)
	//fmt.Println("计算数据长度是 ", lastSize)
	//fmt.Println("真实数据长度是 ", realSize)
	data = make([]byte, lastSize) //先读取头
	size, err = rw.Read(data)
	if err != nil {
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: nil,
		}
		return
	}
	errAndData <- &reqeustErrorAndData{
		err:  err,
		data: data,
	}
}
func timeoutMonitor(address, operator string, timeout int, params interface{}, errAndData chan *reqeustErrorAndData) {
	fmt.Println("启动超时监控启动")
	<-time.After(time.Duration(timeout) * time.Second)
	errAndData <- &reqeustErrorAndData{
		err:  NewTimeoutError(address, operator, timeout, params),
		data: nil,
	}
	fmt.Println("超时完成")
}

//写完上面的head 10字节
// 写 总长度
// 写请求id
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
