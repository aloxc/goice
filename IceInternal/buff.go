package IceInternal

import (
	"bufio"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"sync/atomic"
)

type IceBuffer struct {
	*bufio.ReadWriter
}
var requestId int32 = 2
func NewIceBuff(rw *bufio.ReadWriter) *IceBuffer {
	return &IceBuffer{
		rw,
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
	this.WriteStr(operator)//29+1+8=38
}
//这个值是总数据长度减去设置完context后的长度 TODO
func (this *IceBuffer) WriteRealSize(siz int) {
	this.Write(utils.IntToBytes(siz))
}
func (this *IceBuffer) WriteMode(mode byte) {
	this.WriteByte(mode)//38+1=39
}
func (this *IceBuffer) WriteSize(v int) {
	if (v > 254) { //如果大于254，就写负一，然后跟上四个字节（int长度）的长度，
		//int8 := -1
		//this.
		bytes := utils.IntToBytes(v)
		this.WriteString("")
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
		this.WriteString(str)//	写字符串
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
	this.WriteStringMap(context)//39+1=40 TODO 从这以后就是就是修正长度
}

func (this*IceBuffer)WriteContentLength(length int){
	this.Write(utils.IntToBytes(length))//10+4=14
}
func (this*IceBuffer)WriteEncodingVersion(){
	encoding := GetDefaultEncodingVersion()
	this.WriteByte(encoding.Major)//encoding major 44+1=45
	this.WriteByte(encoding.Minor)//encoding minor 45+1=46
}
func (this*IceBuffer)WriteRequestId(){
	atomic.AddInt32(&requestId,1)
	this.Write(utils.IntToBytes(int(requestId)))//14+4=18
}
func (this*IceBuffer)WriteIdentity(identity *ice.Identity){
	this.WriteStr(identity.Name)//18+1+8=27
	this.WriteStr(identity.Category)//27+1=28
}
func (this*IceBuffer)WriteHead(){
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
	this.Write(requestHead)//10字节
}
