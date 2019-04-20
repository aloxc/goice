package ice

import (
	"bufio"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/utils"
	"github.com/siddontang/go-log/log"
	"reflect"
)

type IceBuffer struct {
	*bufio.ReadWriter
	Out *Outgoing
	In  *Incoming
}

func NewIceBuff(rw *bufio.ReadWriter) *IceBuffer {
	return &IceBuffer{
		ReadWriter: rw,
	}
}

func (this *IceBuffer) WriteFacet(facet string) {
	if len(facet) == 0 {
		this.WriteByte(0) //28+1=29
	} else {
		facets := []string{facet}
		this.WriteStringArray(facets)
	}
}
func (this *IceBuffer) WriteOperator(operator string) {
	this.WriteStr(operator) //29+1+8=38
}

func (this *IceBuffer) WriteBoolArray(arr []bool) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.BoolToBytes(v))
	}
}

func (this *IceBuffer) WriteInt8Array(arr []int8) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.Int8ToBytes(v))
	}
}

func (this *IceBuffer) WriteInt16Array(arr []int16) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.Int16ToBytes(v))
	}
}

func (this *IceBuffer) WriteIntArray(arr []int) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.IntToBytes(v))
	}
}

func (this *IceBuffer) WriteInt32Array(arr []int32) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.Int32ToBytes(v))
	}
}

func (this *IceBuffer) WriteInt64Array(arr []int64) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.Int64ToBytes(v))
	}
}

func (this *IceBuffer) WriteFloat32Array(arr []float32) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.Float32ToBytes(v))
	}
}

func (this *IceBuffer) WriteFloat64Array(arr []float64) {
	this.WriteSize(len(arr))
	for _, v := range arr {
		this.Write(utils.Float64ToBytes(v))
	}
}

//这个值是总数据长度减去设置完context后的长度
func (this *IceBuffer) WriteRealSize(siz int) {
	this.Write(utils.IntToBytes(siz))
}
func (this *IceBuffer) WriteMode(mode byte) {
	this.WriteByte(mode) //38+1=39
}
func (this *IceBuffer) WriteSize(v int) {
	if v > 254 { //如果大于254，就写负一，然后跟上四个字节（int长度）的长度，
		//int8 := -1
		//this.
		bytes := utils.IntToBytes(v)
		this.WriteByte(255)
		this.Write(bytes)
	} else { //否则扩容一个，写一字节的长度
		uv := uint8(v)
		this.WriteByte(uv)
	}
}
func (this *IceBuffer) WriteStr(str string) {
	len := len(str)
	this.WriteSize(len)
	if len > 0 {
		this.WriteString(str) //	写字符串
	}
}
func (this *IceBuffer) WriteMap(context map[string]string) {
	len := len(context)
	this.WriteSize(len)
	if len > 0 {
		for k, v := range context {
			this.WriteStr(k)
			this.WriteStr(v)
		}
	}
}
func (this *IceBuffer) WriteStringArray(arr []string) {
	len := len(arr)
	this.WriteSize(len)
	if len != 0 {
		for _, v := range arr {
			this.WriteStr(v)
		}
	}
}
func (this *IceBuffer) WriteStringMap(m map[string]string) {
	len := len(m)
	this.WriteSize(len)
	if len != 0 {
		for k, v := range m {
			this.WriteStr(k)
			this.WriteStr(v)
		}
	}
}

func (this *IceBuffer) WriteContext(context map[string]string) {
	this.WriteStringMap(context) //39+1=40 从这以后就是就是修正长度
}

//这个是总长度
func (this *IceBuffer) WriteTotalSize(siz int) {
	this.Write(utils.IntToBytes(siz)) //10+4=14
}
func (this *IceBuffer) WriteEncodingVersion(encoding *EncodingVersion) {
	this.WriteByte(encoding.Major) //encoding major 44+1=45
	this.WriteByte(encoding.Minor) //encoding minor 45+1=46
}
func (this *IceBuffer) WriteRequestId(requestId int) {
	this.Write(utils.IntToBytes(requestId)) //14+4=18
}
func (this *IceBuffer) WriteIdentity(identity *Identity) {
	this.WriteStr(identity.name)     //18+1+8=27
	this.WriteStr(identity.category) //27+1=28
}
func (this *IceBuffer) WriteHead() {
	var magic = []byte{0x49, 0x63, 0x65, 0x50}
	var msgType byte = 0

	requestHead := []byte{magic[0],
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
	this.Write(requestHead) //10字节
}

//每次请求要计算请求长度
//TODO 多参还没看怎么处理的，当前就是一个参数的
func (this *IceBuffer) Prepare(identity *Identity, facet, operator string, params interface{}, context map[string]string) (total, real int) {
	total = 0
	total += 10                 //head
	total += 4                  //total(int = 4)
	total += 4                  //requestId(int = 4)
	total += 1                  //identity.Name的长度(byte = 1)
	total += len(identity.name) //identity.Name本身(len(identity.name))
	total += 1                  //identity.Category的长度(byte = 1)
	if len(identity.category) != 0 {
		total += len(identity.category) //identity.Name本身(len(identity.category))
	}
	if len(facet) == 0 {
		total += 1 //facet的长度(byte = 1)
	} else {
		total += 1          //facet数组的长度(byte = 1)
		total += 1          //facet数组[0]的长度(byte = 1)
		total += len(facet) //facet数组[0]的本身(len(facet))
	}
	total += 1             //operator的长度(byte = 1)
	total += len(operator) //operator的本身的长度
	total += 1             //mode的长度(byte = 1)
	if context == nil || len(context) == 0 {
		total += 1 //context的长度(byte = 1)
	} else {
		if len(context) > 254 {
			total += 1 // -1 超过254 就设置个-1和int
			total += 4 //int
		} else {
			total += 1 //param的长度
			for k, v := range context {
				total += 1      // key
				total += 1      // value
				total += len(k) // key 本身
				total += len(v) // value 本身
			}
		}
	}
	end := total
	total += 4 //整形后的数据长度（int = 4 ）
	total += 1 //encoding major
	total += 1 //encoding manor
	log.Info("参数类型", reflect.TypeOf(params))
	if params != nil {
		for _, param := range params.([]interface{}) {
			switch param.(type) {
			case string:
				total += getArraySize(1, len(param.(string)))
			case []string:
				if len(param.([]string)) > 254 {
					total += 1 // -1 超过254 就设置个-1和int
					total += 4 //
				} else {
					total += 1 //param的长度
				}
				for _, sub := range param.([]string) {
					total += getArraySize(1, len(sub))
				}
			case bool:
				total += 1
			case []bool:
				total += getArraySize(1, len(param.([]bool)))
			case int8:
				total += 1
			case []int8:
				total += getArraySize(1, len(param.([]int8)))
			case int16:
				total += 2
			case []int16:
				total += getArraySize(2, len(param.([]int16)))
			case int:
				total += 4
			case []int:
				total += getArraySize(4, len(param.([]int)))
			case int32:
				total += 4
			case []int32:
				total += getArraySize(4, len(param.([]int32)))
			case int64:
				total += 8
			case []int64:
				total += getArraySize(8, len(param.([]int64)))
			case float32:
				total += 4
			case []float32:
				total += getArraySize(4, len(param.([]float32)))
			case float64:
				total += 8
			case []float64:
				total += getArraySize(8, len(param.([]float64)))
			case *Request:
				//fmt.Println("进入到了Request类型计算长度")
				request := param.(*Request)
				total += 1
				//fmt.Println("方法:",request.Method)
				total += len(request.Method)
				total += 1
				for k, v := range request.Params {
					//fmt.Println("参数:",k,v)
					total += 1
					total += len(k)
					total += 1
					total += len(v)
				}

			}
		}
	}

	//log.Info("请求长度 ", total, end)
	return total, total - end
}
func getArraySize(per, length int) int {
	var total = 0
	if length > 254 {
		total += 1 // -1 超过254 就设置个-1和int
		total += 4 //
	} else {
		total += 1 //param的长度
	}
	total += length * per // params本身长度
	return total
}

//连接创建后要发送head请求，计算长度
func PrepareHead(identity *Identity, facet, module string, context map[string]string) (total, real int) {
	total = 0
	total += 10                 //head
	total += 4                  //total(int = 4)
	total += 4                  //requestId(int = 4)
	total += 1                  //identity.Name的长度(byte = 1)
	total += len(identity.name) //identity.Name本身(len(identity.name))
	total += 1                  //identity.Category的长度(byte = 1)
	if len(identity.category) != 0 {
		total += len(identity.category) //identity.Name本身(len(identity.category))
	}
	if len(facet) == 0 {
		total += 1 //facet的长度(byte = 1)
	} else {
		total += 1          //facet数组的长度(byte = 1)
		total += 1          //facet数组[0]的长度(byte = 1)
		total += len(facet) //facet数组[0]的本身(len(facet))
	}

	total += 1                   //operator的长度(byte = 1、len(operator))
	total += len(config.Ice_isA) //operator的本身的长度
	total += 1                   //mode的长度(byte = 1)
	if context == nil || len(context) == 0 {
		total += 1 //context的长度(byte = 1)
	} else {
		if len(context) > 254 {
			total += 1 // -1 超过254 就设置个-1和int
			total += 4 //int
		} else {
			total += 1 //param的长度
			for k, v := range context {
				total += 1      // key
				total += 1      // value
				total += len(k) // key 本身
				total += len(v) // value 本身
			}
		}
	}
	end := total
	//数据整形 java 中 BasicStream.endWriteEncaps方法，大约344行，写此后（39位后的）的数据长度，总数据长度减去39
	total += 4 //整形后的数据长度（int = 4 ）
	total += 1 //encoding major
	total += 1 //encoding manor
	//::module名（slice中定义的）::服务名（slice中定义的interface）
	total += 1                                        // (::module::name)的长度
	total += 2 + len(module) + 2 + len(identity.name) // 模块长度
	//log.Info("请求长度 ", total, end)
	return total, total - end
}
