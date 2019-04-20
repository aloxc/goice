package ice

import (
	"bufio"
	"encoding/json"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/pool"
	"github.com/aloxc/goice/utils"
	"github.com/pkg/errors"
	"github.com/siddontang/go/log"
	"io"
	"net"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Request struct {
	Method string
	Params map[string]string
}

type reqeustErrorAndData struct {
	err  error
	data interface{}
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

var (
	requestId int32 = 2
	connMap         = map[string]pool.Pool{}
	mux       sync.Mutex
)

type IceRequest struct {
	name      string    //对应servers下面的子节点名称
	head      *[]byte   // 10 字节 0 + 10 = 10
	totalSize int       //所有数据长度 4 字节 10 + 4 = 14
	requestId int       //请求id 4 字节 14 + 4 = 18
	Identity  *Identity //该对象标识  1 + xx + 1 + 0 = 10 字节  18 + 2 + xx = 20 + xx
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

//准备把所有设置都放到这个方法中，先Prepare下，然后再调用组装数据的，最后就是执行this.Flush
func (this *IceRequest) DoRequest(responseType ResponseType) (interface{}, error) {
	var timeout int = 5
	atomic.AddInt32(&requestId, 1)
	this.requestId = int(requestId)
	curPool, err := this.getPool(this.name)
	if err != nil { //如果连接失败。则返回。
		log.Error(err)
		return nil, err
	}
	conn, err := curPool.Get()
	if err != nil { //如果连接失败。则返回。
		log.Error(err)
		return nil, err
	}
	defer curPool.Return(conn)
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	var buf = NewIceBuff(rw)
	if this.Context == nil {
		this.Context = make(map[string]string)
	}
	this.Context[string(config.Context_ClientAddr)] = conn.LocalAddr().String() //往后端传客户端地址
	log.Info("参数长度", len(this.Params.([]interface{})))
	//if this.Params != nil {
	//	if len(this.Params.([]interface{})) == 1 {
	//		this.Params = this.Params.([]interface{})[0]
	//	} else if len(this.Params.([]interface{})) == 0 {
	//		this.Params = nil
	//	}
	//}

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
	//this.writeParams(buf)
	log.Info("参数类型", reflect.TypeOf(this.Params))
	//os.Exit(1)
	if this.Params != nil {
		for _, param := range this.Params.([]interface{}) {
			switch param.(type) {
			case string:
				buf.WriteStr(param.(string))
			case []string:
				buf.WriteStringArray(param.([]string))
			case bool:
				buf.Write(utils.BoolToBytes(param.(bool)))
			case []bool:
				buf.WriteBoolArray(param.([]bool))
			case int8:
				buf.Write(utils.Int8ToBytes(param.(int8)))
			case []int8:
				buf.WriteInt8Array(param.([]int8))
			case int16:
				buf.Write(utils.Int16ToBytes(param.(int16)))
			case []int16:
				buf.WriteInt16Array(param.([]int16))
			case int:
				buf.Write(utils.IntToBytes(param.(int)))
			case []int:
				buf.WriteIntArray(param.([]int))
			case int32:
				buf.Write(utils.Int32ToBytes(param.(int32)))
			case []int32:
				buf.WriteInt32Array(param.([]int32))
			case int64:
				buf.Write(utils.Int64ToBytes(param.(int64)))
			case []int64:
				buf.WriteInt64Array(param.([]int64))
			case float32:
				buf.Write(utils.Float32ToBytes(param.(float32)))
			case []float32:
				buf.WriteFloat32Array(param.([]float32))
			case float64:
				buf.Write(utils.Float64ToBytes(param.(float64)))
			case []float64:
				buf.WriteFloat64Array(param.([]float64))
			case *Request:
				request := param.(*Request)
				buf.WriteStr(request.Method)
				buf.WriteStringMap(request.Params)
			}
		}
	}

	buf.Flush()
	//var timeoutCh chan int

	errAndData := make(chan *reqeustErrorAndData)
	go request(conn.RemoteAddr().String(), rw, responseType, this.Operator, this.Params, errAndData)
	go requestTimeoutMonitor(conn.RemoteAddr().String(), this.Operator, timeout, this.Params, errAndData)
	ed := <-errAndData
	return ed.data, ed.err

}
func (this *IceRequest) writeParams(buf *IceBuffer) {
	if this.Params == nil {
		return
	}
	tp := reflect.TypeOf(this.Params).String()
	log.Info(tp, "==")
	if strings.HasPrefix("*[] ", tp) || strings.HasPrefix("[] ", tp) {

	}
}

func (this *IceRequest) writeOneParam(buf *IceBuffer) {
}

func NewIceRequest(name string, mode OperatorMode, operator string, context map[string]string, params ...interface{}) *IceRequest {
	return &IceRequest{
		name:            name,
		head:            GetHead(),
		Operator:        operator,
		OperatorMode:    mode,
		encodingVersion: GetDefaultEncodingVersion(),
		Params:          params,
		Identity:        GetIdentity(config.ConfigMap[name][config.IdentityName].(string), ""),
		Context:         context,
	}
}

//初始化连接池
func (this *IceRequest) getPool(name string) (pool.Pool, error) {
	var curPool pool.Pool
	var ok bool
	if curPool, ok = connMap[name]; !ok {
		//mux.Lock()
		if curPool, ok = connMap[name]; !ok { //双检查锁
			var curConn *net.Conn
			curPool, err := pool.NewGPool(
				&pool.PoolConfig{
					Factory: func() (net.Conn, error) {
						conn, err := net.DialTimeout("tcp4", config.ConfigMap[name][config.Address].(string),
							time.Duration(config.ConfigMap[name][config.ConnectTimeout].(int))*time.Second)

						if err == nil {
							curConn = &conn
							InitConnection(this.Identity, this.name, curConn) //连接后一定要向服务器发送一条head请求
						}
						//mux.Unlock()

						return conn, err
					},
					MaxCap:  config.ConfigMap[name]["MaxClientSize"].(int),
					InitCap: 1,
				})
			if err != nil {
				//mux.Unlock()
				return nil, errors.New("连接池异常")
			}
			curPool.Return(*curConn) //这个连接发起过head请求，也要先还回去。
			connMap[name] = curPool
		}
		//mux.Unlock()
	}
	curPool = connMap[name]
	return curPool, nil

	//直接使用连接的代码
	//var conn, err = Connect("tcp4", address, timeout)
	//if err != nil { //如果连接失败。则返回。
	//	log.Error(err)
	//	return nil, err
	//}
}
func request(address string, rw io.ReadWriter, responseType ResponseType, operator string, params interface{}, errAndData chan *reqeustErrorAndData) {
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
	//pmj = head[4]
	//pmn = head[5]
	//emj = head[6]
	//emn = head[7]
	//rmsg = head[8]
	//zip = head[9]
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
	var replyStatus = head[18]
	switch replyStatus {
	case ReplyUserException:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewUserError(address, operator, string(data), params),
			data: nil,
		}
	case ReplyObjectNotExist:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewObjectNotExistsError(address, operator, string(data), params),
			data: nil,
		}
	case ReplyFacetNotExist:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewFacetNotExistsError(address, operator, string(data), params),
			data: nil,
		}
		return
	case ReplyOperationNotExist:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewOperatorNotExistsError(address, operator, string(data), params),
			data: nil,
		}
		return
	case ReplyUnknownLocalException:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewIceServerError(address, operator, string(data), params),
			data: nil,
		}
	case ReplyUnknownUserException:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		errAndData <- &reqeustErrorAndData{
			err:  NewUserError(address, operator, string(data), params),
			data: nil,
		}
	case ReplyUnknownException: //用户异常
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			errAndData <- &reqeustErrorAndData{
				err:  err,
				data: nil,
			}
			return
		}
		data = append([]byte{head[24]}, data...)
		userUnknownError := NewUserUnknownError(address, operator, string(data), params)
		errAndData <- &reqeustErrorAndData{
			err:  userUnknownError,
			data: nil,
		}
		return

	}
	lastSize = utils.BytesToInt(head[19:23])
	//log.Info("整形后面的数据长度（包括整形4字节） ", lastSize)
	//_encodingMajor := head[23]
	//_encodingMinor := head[24]
	lastSize = lastSize - 4 - 1 - 1 //4:整形后面包括整形长度，1：主编码版本 ，1：副编码版本
	//fmt.Printf("编码版本major = %d,minor = %d\n", _encodingMajor, _encodingMinor)
	//log.Info("最终数据长度及数据 的长度", lastSize)
	//TODO 这些还要处理数组问题
	if responseType == ResponseType_String {
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		_, offset := readSize(data)
		log.Info("string结果", string(data[offset:]))
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: data[offset:],
		}
		return
	} else if responseType == ResponseType_String_Array {
		log.Info("剩余", lastSize)
		data := make([]byte, lastSize)

		size, err = rw.Read(data)
		var arr = readStringArray(data)
		log.Info("字符串数组", arr)
		//for i,v:=range arr {
		//	log.Infof("[%d] = %d，[%s]",i, len(v),v)
		//}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
	} else if responseType == ResponseType_Bool { //通过
		data := make([]byte, 1)
		size, err = rw.Read(data)
		var re = utils.BytesToBool(data)
		log.Info("bool结果 ", re)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: re,
		}
		return
	} else if responseType == ResponseType_Bool_Array { //通过
		log.Info(lastSize)
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]bool, size)
		for i := 0; i < lastSize-1; i++ {
			arr[i] = utils.BytesToBool(data[i+offset : i+1+offset])
		}
		log.Info("boolArray结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Int8 { //测试通过
		data := make([]byte, 1)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToInt8(data),
		}
		return
	} else if responseType == ResponseType_Int8_Array { //测试通过

		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		log.Info("读取int8Array", lastSize, size, offset)
		var arr = make([]int8, size)
		for i := 0; i < lastSize-1; i++ {
			log.Info("i =", i)
			arr[i] = utils.BytesToInt8(data[i+offset : i+1+offset])
		}
		log.Info("int8Array结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Int16 { //通过
		data := make([]byte, 2)
		size, err = rw.Read(data)
		log.Info("int16=", utils.BytesToInt16(data))
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToInt16(data),
		}
		return
	} else if responseType == ResponseType_Int16_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		log.Info("读取int16Array", lastSize, size, offset)
		var arr = make([]int16, size)
		for i := 0; i < size; i++ {
			log.Info("i =", i)
			arr[i] = utils.BytesToInt16(data[i*2+offset : (i+1)*2+offset])
		}
		log.Info("int16Array结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Int { //通过
		data := make([]byte, 4)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToInt(data),
		}
		return
	} else if responseType == ResponseType_Int_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		log.Info("读取intArray", lastSize, size, offset)
		var arr = make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt(data[i*4+offset : (i+1)*4+offset])
		}
		log.Info("intArray结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Int64 {
		data := make([]byte, 8)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToInt64(data),
		}
		return
	} else if responseType == ResponseType_Int64_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		log.Info("读取int64Array", lastSize, size, offset)
		var arr = make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt(data[i*8+offset : (i+1)*8+offset])
		}
		log.Info("int64Array结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Float32 {
		data := make([]byte, 4)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToFloat32(data),
		}
		return
	} else if responseType == ResponseType_Float32_Array {
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		log.Info("读取float32Array", lastSize, size, offset)
		var arr = make([]float32, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToFloat32(data[i*4+offset : (i+1)*4+offset])
		}
		log.Info("float32Array结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Float64 {
		log.Info("lastSize = ", lastSize)
		data := make([]byte, 8)
		size, err = rw.Read(data)
		log.Info("float64:", utils.BytesToFloat64(data))
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToFloat64(data),
		}
		return
	} else if responseType == ResponseType_Float64_Array {
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		log.Info("读取float64Array", lastSize, size, offset)
		var arr = make([]float64, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToFloat64(data[i*8+offset : (i+1)*8+offset])
		}
		log.Info("float64Array结果", arr)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	}
	//sizeDefine := 0
	////realSize := 0
	////字符串的
	//if lastSize <= 255 {
	//	//读取一个字节
	//	sizeDefine = 1
	//	dataSizeData := make([]byte, sizeDefine)
	//	size, err = rw.Read(dataSizeData)
	//	if err != nil {
	//		errAndData <- &reqeustErrorAndData{
	//			err:  err,
	//			data: nil,
	//		}
	//		return
	//	}
	//	//realSize = int(dataSizeData[0])
	//} else {
	//	//读取一个字节-1，跟着4个字节的数据长度（int）
	//	sizeDefine = 5
	//	dataSizeData := make([]byte, sizeDefine)
	//	size, err = rw.Read(dataSizeData)
	//	if err != nil {
	//		errAndData <- &reqeustErrorAndData{
	//			err:  err,
	//			data: nil,
	//		}
	//		return
	//	}
	//	//realSize = utils.BytesToInt(dataSizeData[1:])
	//	//log.Info("长度超过254，读取下这个 -1 是什么%d", dataSizeData[0])
	//}
	////log.Info("lastSize=",lastSize)
	//lastSize = lastSize - sizeDefine
	////log.Info("sizeDefine=" ,sizeDefine)
	////log.Info("计算数据长度是 ", lastSize)
	////log.Info("真实数据长度是 ", realSize)
	//data = make([]byte, lastSize) //先读取头
	//size, err = rw.Read(data)
	//if err != nil {
	//	errAndData <- &reqeustErrorAndData{
	//		err:  err,
	//		data: nil,
	//	}
	//	return
	//}
	errAndData <- &reqeustErrorAndData{
		err:  err,
		data: data,
	}
}

//请求超时monitor
func requestTimeoutMonitor(address, operator string, timeout int, params interface{}, errAndData chan *reqeustErrorAndData) {
	log.Info("启动超时监控启动")
	<-time.After(time.Duration(timeout) * time.Second)
	errAndData <- &reqeustErrorAndData{
		err:  NewTimeoutError(address, operator, timeout, params),
		data: nil,
	}
	log.Info("超时完成")
}
func readString(data []byte, offset, size int) (str string) {
	str = string(data[offset : offset+size])
	return
}

func readIntArray(data []byte) (arr []int32) {
	size, offset := readSize(data)
	arr = make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = utils.BytesToInt32(data[i*4+offset : (i+1)*4+offset])
	}
	return
}

//读取数据长度或者数组长度
func readSize(data []byte) (size, offset int) {
	offset = 0
	if data[0] == 255 {
		offset = 5
		size = utils.BytesToInt(data[1:5])
	} else {
		offset = 1
		size = int(data[0])
	}
	return
}
func readStringArray(data []byte) (arr []string) {
	if data[0] == 0 {
		return
	}
	arrSize, offset := readSize(data)
	arr = make([]string, arrSize)
	step := 0
	size := 0
	for i := 0; i < arrSize; i++ {
		if data[offset] == 255 {
			size = utils.BytesToInt(data[offset+1 : offset+5])
			step = 1 + 4
		} else {
			size = int(data[offset])
			step = 1
		}
		arr[i] = string(data[offset+step : offset+size+step])
		offset += step + size
	}
	return
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
