package IceInternal

import (
	"bufio"
)

type IceBuffer struct {
	*bufio.ReadWriter
}

func NewIceBuff(rw *bufio.ReadWriter) *IceBuffer {
	return &IceBuffer{
		rw,
	}

}
func (this *IceBuffer) WriteSize(v int) {
	if (v > 254) { //如果大于254，就写负一，然后跟上四个字节（int长度）的长度，
		//int8 := -1
		//this.
		this.WriteString("")
		this.WriteByte(0)
		this.WriteByte(0)
		this.WriteByte(0)
		this.WriteByte(0)
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

