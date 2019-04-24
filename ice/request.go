package ice

import (
	"bufio"
	"encoding/json"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/utils"
	"github.com/siddontang/go/log"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)
type Request struct {
	Method string
	Params map[string]string
}

var showResult = false

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
	requestId   int32 = 2
	connPoolMap       = map[string]Pool{}
	mux         sync.Mutex
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
	Operation       string            //操作名称，也就是要调用的方法名称 yy 字节 21 + xx + yy
	OperationMode                     // 1 字节 1 + 21 + xx + yy = 22 + xx + yy
	realSize        int               // 4 字节 4 + 22 + xx + yy = 26 + xx + yy
	Context         map[string]string //调用上下文 zz 字节
	encodingVersion *EncodingVersion  // 2 字节
	Params          interface{}
	OperateTimeout int
}

//准备把所有设置都放到这个方法中，先Prepare下，然后再调用组装数据的，最后就是执行this.Flush
func (this *IceRequest) DoRequest(responseType ResponseType) (interface{}, error) {
	//var timeout int = 5
	atomic.AddInt32(&requestId, 1)
	this.requestId = int(requestId)
	curPool, err := this.getPool(this.name)
	if err != nil { //如果连接失败。则返回。
		log.Error(err)
		return nil, err
	}
	//log.Info("剩余", len(curPool.freeConns),curPool,curPool.freeConns)
	conn, err := curPool.Get()
	defer func() {
		curPool.Return(conn)
		//log.Info("归还中", len(curPool.freeConns))
		connPoolMap[this.name] = curPool
		//tPool = curPool
	}()
	//直接使用连接的代码
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	var buf = NewIceBuff(rw)
	if this.Context == nil {
		this.Context = make(map[string]string)
	}
	this.Context[string(config.Context_ClientAddr)] = conn.LocalAddr().String() //往后端传客户端地址
	//log.Info("参数个数", len(this.Params.([]interface{})))

	total, real := buf.Prepare(this.Identity, this.Facet, this.Operation, this.Params, this.Context)
	buf.Write(*this.head)
	buf.WriteTotalSize(total)
	buf.WriteRequestId(this.requestId)
	buf.WriteIdentity(this.Identity)
	buf.WriteFacet(this.Facet)
	buf.WriteOperator(this.Operation)
	buf.WriteByte(byte(this.OperationMode))
	buf.WriteContext(this.Context)
	buf.WriteRealSize(real)
	buf.WriteEncodingVersion(this.encodingVersion)
	//this.writeParams(buf)
	//log.Info("参数类型", reflect.TypeOf(this.Params))
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

	errAndData := make(chan *reqeustErrorAndData)
	go doResult(conn.RemoteAddr().String(), rw, responseType, this.Operation, this.Params, errAndData)
	//return doResultDirect(conn.RemoteAddr().String(), rw, responseType, this.Operation, this.Params)
	go requestTimeoutMonitor(conn.RemoteAddr().String(), this.Operation, this.OperateTimeout, this.Params, errAndData)
	ed := <-errAndData
	return ed.data, ed.err

}
func NewIceRequest(name string, mode OperationMode, operator string, context map[string]string, params ...interface{}) *IceRequest {
	return &IceRequest{
		name:            name,
		head:            GetHead(),
		Operation:       operator,
		OperationMode:   mode,
		encodingVersion: GetDefaultEncodingVersion(),
		Params:          params,
		Identity:        GetIdentity(config.ConfigMap[name][config.IdentityName].(string), ""),
		Context:         context,
		OperateTimeout:config.ConfigMap[name][config.OperateTimeout].(int),
	}
}

var count int32 = 0
//初始化连接池
func (this *IceRequest) getPool(name string) (Pool, error) {
	var curPool Pool
	var ok bool
	curPool, ok = connPoolMap[name]
	//log.Info(ok)
	if !ok {
		if curPool, ok = connPoolMap[name]; !ok { //双检查锁
			curPool := Pool{
				Network: "tcp4",
				Address: config.ConfigMap[name][config.Address].(string),
				NewConnHook: &IceNewConnHook{
					Identity: GetIdentity(config.ConfigMap[name][config.Name].(string), ""),
					Name:     config.ConfigMap[name][config.Name].(string),
				},
				MaxConn:     config.ConfigMap[name]["MaxClientSize"].(int),
				MaxLifetime: time.Duration(config.ConfigMap[name][config.MaxIdleTime].(int))*time.Second ,
			}
			connPoolMap[name] = curPool
		}
	}
	curPool = connPoolMap[name]
	return curPool, nil
}

func doResult(address string, rw io.ReadWriter, responseType ResponseType, operator string, params interface{}, errAndData chan *reqeustErrorAndData) {
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
		if showResult {
			log.Info("string结果", string(data[offset:]))
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: string(data[offset:]),
		}
		return
	} else if responseType == ResponseType_String_Array {
		data := make([]byte, lastSize)

		size, err = rw.Read(data)
		var arr = readStringArray(data)
		if showResult {
			log.Info("字符串数组", arr)
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
	} else if responseType == ResponseType_Bool { //通过
		data := make([]byte, 1)
		size, err = rw.Read(data)
		var re = utils.BytesToBool(data)
		if showResult {
			log.Info("bool结果 ", re)
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToBool(data),
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
		if showResult {
			log.Info("boolArray结果", arr)
		}
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
		var arr = make([]int8, size)
		for i := 0; i < lastSize-1; i++ {
			arr[i] = utils.BytesToInt8(data[i+offset : i+1+offset])
		}
		if showResult {
			log.Info("int8Array结果", arr)
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Int16 { //通过
		data := make([]byte, 2)
		size, err = rw.Read(data)
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToInt16(data),
		}
		return
	} else if responseType == ResponseType_Int16_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]int16, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt16(data[i*2+offset : (i+1)*2+offset])
		}
		if showResult {
			log.Info("int16Array结果", arr)
		}
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
		var arr = make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt(data[i*4+offset : (i+1)*4+offset])
		}
		if showResult {
			log.Info("intArray结果", arr)
		}
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
		var arr = make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt(data[i*8+offset : (i+1)*8+offset])
		}
		if showResult {
			log.Info("int64Array结果", arr)
		}
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
		if showResult {
			log.Info("读取float32Array", lastSize, size, offset)
		}
		var arr = make([]float32, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToFloat32(data[i*4+offset : (i+1)*4+offset])
		}
		if showResult {
			log.Info("float32Array结果", arr)
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	} else if responseType == ResponseType_Float64 {
		data := make([]byte, 8)
		size, err = rw.Read(data)
		if showResult {
			log.Info("float64:", utils.BytesToFloat64(data))
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: utils.BytesToFloat64(data),
		}
		return
	} else if responseType == ResponseType_Float64_Array {
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]float64, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToFloat64(data[i*8+offset : (i+1)*8+offset])
		}
		if showResult {
			log.Info("float64Array结果", arr)
		}
		errAndData <- &reqeustErrorAndData{
			err:  err,
			data: arr,
		}
		return
	}
	errAndData <- &reqeustErrorAndData{
		err:  err,
		data: data,
	}
}
func doResultDirect(address string, rw io.ReadWriter, responseType ResponseType, operator string, params interface{}) (interface{}, error) {
	var size, lastSize int
	var head []byte
	head = make([]byte, 25) //先读取头
	size, err := rw.Read(head)
	if err != nil {
		return nil, err
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
		return nil, err
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
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		return nil, NewUserError(address, operator, string(data), params)
	case ReplyObjectNotExist:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		return nil, NewObjectNotExistsError(address, operator, string(data), params)
	case ReplyFacetNotExist:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		return nil, NewFacetNotExistsError(address, operator, string(data), params)
	case ReplyOperationNotExist:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		return nil, NewOperatorNotExistsError(address, operator, string(data), params)
	case ReplyUnknownLocalException:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		return nil, NewIceServerError(address, operator, string(data), params)
	case ReplyUnknownUserException:
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		return nil, NewUserError(address, operator, string(data), params)
	case ReplyUnknownException: //用户异常
		lastSize = utils.BytesToInt(head[20:24])
		data := make([]byte, lastSize) //读取用户异常信息
		size, err = rw.Read(data)
		if err != nil {
			return nil, err
		}
		data = append([]byte{head[24]}, data...)
		userUnknownError := NewUserUnknownError(address, operator, string(data), params)
		return nil, userUnknownError
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
		if showResult {
			log.Info("string结果", string(data[offset:]))
		}
		return string(data[offset:]), err
	} else if responseType == ResponseType_String_Array {
		data := make([]byte, lastSize)

		size, err = rw.Read(data)
		var arr = readStringArray(data)
		if showResult {
			log.Info("字符串数组", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Bool { //通过
		data := make([]byte, 1)
		size, err = rw.Read(data)
		var re = utils.BytesToBool(data)
		if showResult {
			log.Info("bool结果 ", re)
		}
		return utils.BytesToBool(data), err
	} else if responseType == ResponseType_Bool_Array { //通过
		log.Info(lastSize)
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]bool, size)
		for i := 0; i < lastSize-1; i++ {
			arr[i] = utils.BytesToBool(data[i+offset : i+1+offset])
		}
		if showResult {
			log.Info("boolArray结果", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Int8 { //测试通过
		data := make([]byte, 1)
		size, err = rw.Read(data)
		return utils.BytesToInt8(data), err
	} else if responseType == ResponseType_Int8_Array { //测试通过

		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]int8, size)
		for i := 0; i < lastSize-1; i++ {
			arr[i] = utils.BytesToInt8(data[i+offset : i+1+offset])
		}
		if showResult {
			log.Info("int8Array结果", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Int16 { //通过
		data := make([]byte, 2)
		size, err = rw.Read(data)
		return utils.BytesToInt16(data), err
	} else if responseType == ResponseType_Int16_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]int16, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt16(data[i*2+offset : (i+1)*2+offset])
		}
		if showResult {
			log.Info("int16Array结果", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Int { //通过
		data := make([]byte, 4)
		size, err = rw.Read(data)
		return utils.BytesToInt(data), err
	} else if responseType == ResponseType_Int_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt(data[i*4+offset : (i+1)*4+offset])
		}
		if showResult {
			log.Info("intArray结果", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Int64 {
		data := make([]byte, 8)
		size, err = rw.Read(data)
		return utils.BytesToInt64(data), err
	} else if responseType == ResponseType_Int64_Array { //通过
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToInt(data[i*8+offset : (i+1)*8+offset])
		}
		if showResult {
			log.Info("int64Array结果", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Float32 {
		data := make([]byte, 4)
		size, err = rw.Read(data)
		return utils.BytesToFloat32(data), err
	} else if responseType == ResponseType_Float32_Array {
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		if showResult {
			log.Info("读取float32Array", lastSize, size, offset)
		}
		var arr = make([]float32, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToFloat32(data[i*4+offset : (i+1)*4+offset])
		}
		if showResult {
			log.Info("float32Array结果", arr)
		}
		return arr, err
	} else if responseType == ResponseType_Float64 {
		data := make([]byte, 8)
		size, err = rw.Read(data)
		if showResult {
			log.Info("float64:", utils.BytesToFloat64(data))
		}
		return utils.BytesToFloat64(data), err
	} else if responseType == ResponseType_Float64_Array {
		data := make([]byte, lastSize)
		size, err = rw.Read(data)
		size, offset := readSize(data)
		var arr = make([]float64, size)
		for i := 0; i < size; i++ {
			arr[i] = utils.BytesToFloat64(data[i*8+offset : (i+1)*8+offset])
		}
		if showResult {
			log.Info("float64Array结果", arr)
		}
		return arr, err
	}
	return nil, nil
}

//请求超时monitor
func requestTimeoutMonitor(address, operator string, timeout int, params interface{}, errAndData chan *reqeustErrorAndData) {
	if showResult {
		log.Info("启动超时监控启动")
	}
	<-time.After(time.Duration(timeout) * time.Second)
	errAndData <- &reqeustErrorAndData{
		err:  NewTimeoutError(address, operator, timeout, params),
		data: nil,
	}
	log.Info("超时完成")
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
type IceNewConnHook struct {
	Identity *Identity
	Name string
}
func (this*IceNewConnHook)hook(conn *net.Conn) error{
	log.Info("开始执行hook = ", atomic.AddInt32(&count, 1))
	var facet string
	rw := bufio.NewReadWriter(bufio.NewReader(*conn), bufio.NewWriter(*conn))
	var buf = NewIceBuff(rw)

	total, real := PrepareHead(this.Identity, "", config.ConfigMap[this.Name][config.Module].(string), nil)

	var context map[string]string
	var head, data []byte
	var err error
	buf.Write(*GetConnHead())                     // 10字节
	buf.Write(utils.IntToBytes(total))                //size 10 +4 = 14
	buf.Write(utils.IntToBytes(1))                    //requestId 14+4=18
	buf.WriteStr(this.Identity.GetIdentityName())     //18+1+8=27
	buf.WriteStr(this.Identity.GetIdentityCategory()) //27+1=28

	if len(facet) == 0 {
		buf.WriteByte(0) //28+1=29
	} else {
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}
	buf.WriteStr(string(config.Ice_isA))             //29+1+7=37
	buf.WriteByte(byte(OperatorModeNonmutating)) //37+1=38
	context = make(map[string]string)
	buf.WriteStringMap(context) //38+1=39
	//数据整形 java 中 BasicStream.endWriteEncaps方法，大约344行，写此后（39位后的）的数据长度，总数据长度减去39
	buf.Write(utils.IntToBytes(real)) //修正数据长度，是总长度减去写完context后的长度 39 +4 = 43

	buf.WriteByte(1) //encoding major 43+1=44
	buf.WriteByte(1) //encoding minor 44+1=45

	buf.WriteStr("::" + config.ConfigMap[this.Name][config.Module].(string) + "::" + config.ConfigMap[this.Name][config.Name].(string)) //45+1+23=69
	buf.Flush()
	head = make([]byte, 14) //先读取头
	_, err = rw.Read(head)
	if err != nil {
		return err
	}
	data = make([]byte, 26) //连接的时候
	_, err = rw.Read(data)
	return err
}