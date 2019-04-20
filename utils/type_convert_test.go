package utils

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestBytesToInt(t *testing.T) {
	var b byte = 250
	var i byte = 0
	for i = 0; i < 10; i++ {
		b = b + i
		fmt.Println(b)
	}
	var i8 int8 = -1
	//bytes := utils.IntToBytes(i8)
	fmt.Println(reflect.TypeOf(byte(i8)))
	fmt.Println(unsafe.Sizeof(byte(i8)))
	var f float32 = 234.43
	fmt.Println(f)
	bytes := Float32ToByte(f)
	fmt.Println(ByteToFloat32(bytes))

	toBytes := BoolToBytes(false)
	fmt.Println(len(toBytes))
}
