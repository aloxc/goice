package main

import (
	"fmt"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"testing"
)

func TestGoiceSayHi(t *testing.T) {
	//goice.sayHi() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "sayHi", nil, "")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		if showResult {
			fmt.Println("请求结果", string(result))
		}
	}
}
func TestGoiceSayHiUsingContext(t *testing.T) {
	//goice.sayHi(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "sayHi", context, "")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceSayHello(t *testing.T) {
	//hello.sayHello(string) 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "sayHello", nil, "aaa")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceSayHelloUsingContext(t *testing.T) {
	//goice.sayHi(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "sayHello", context, "aaa")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceVoid(t *testing.T) {
	//goice.vvoid() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "vvoid", nil, "")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceVoidUsingContext(t *testing.T) {
	//goice.vvoid(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "vvoid", context, "")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceVoidTo(t *testing.T) {
	//goice.vvoid(string) 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "vvoidTo", nil, "aaa")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceVoidToUsingContext(t *testing.T) {
	//goice.vvoid(string,context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "vvoidTo", context, "aaa")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceGetAge(t *testing.T) {
	//goice.getAge() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getAge", nil, 123321)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		if showResult {
			fmt.Println("请求结果", utils.BytesToInt(result))
		}
	}
}
func TestGoiceGetAgeUsingContext(t *testing.T) {
	//goice.getAge(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getAge", context, 123)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}

func TestGoiceGetAges(t *testing.T) {
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getAges", nil, 123321)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}
func TestGoiceGetAgesUsingContext(t *testing.T) {
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getAges", nil, 123321)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}
func TestGoiceGetBool(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getBool", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetBoolUsingContext(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getBool", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetBoolFrom(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getBoolFrom", nil, false)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetBoolFromUsingContext(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getBoolFrom", nil, true)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetColor(t *testing.T) {
	//通过，但是需要自己把int换成枚举
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getColor", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}
func TestGoiceGetColorUsingContext(t *testing.T) {
	//通过，但是需要自己把int换成枚举
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getColor", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}
func TestGoiceGetColorFrom(t *testing.T) {
	//通过，传枚举需要使用int8 或者int16
	var i8 int8 = 1
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getColorFrom", nil, i8)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}
func TestGoiceGetColorFromUsingContext(t *testing.T) {
	//通过，传枚举需要使用int8 或者int16
	var i8 int8 = 3
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getColorFrom", nil, i8)
	result, err := request.DoRequest(ice.ResponseType_Int)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt(result))
	}
}
func TestGoiceGetByte(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功
	var i8 int8 = 101
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getByte", nil, i8)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt8(result))
	}
}
func TestGoiceGetByteUsingContext(t *testing.T) {
	//goice.getShort(int16,context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	var i8 int8 = 73
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getByte", context, i8)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt8(result))
	}
}
func TestGoiceGetShort(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功
	//var i16 int16 = 321
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getShort", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt16(result))
	}
}
func TestGoiceGetShortUsingContext(t *testing.T) {
	//goice.getShort(int16,context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	//var i16 int16 = 321
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getShort", context, nil)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt16(result))
	}
}
func TestGoiceGetShortFrom(t *testing.T) {
	//goice.getShortFrom(context) 测试成功
	var i16 int16 = 321
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getShortFrom", nil, i16)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt16(result))
	}
}
func TestGoiceGetShortFromUsingContext(t *testing.T) {
	//goice.getShortFrom(int16,context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	var i16 int16 = 110
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getShortFrom", nil, i16)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt16(result))
	}
}
func TestGoiceGetLong(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getLong", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt64(result))
	}
}
func TestGoiceGetLongUsingContext(t *testing.T) {
	//通过
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getLong", context, nil)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt64(result))
	}
}
func TestGoiceGetLongFrom(t *testing.T) {
	//通过
	var i64 int64 = 922337203685477581
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getLongFrom", nil, i64)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt64(result))
	}
}
func TestGoiceGetLongFromUsingContext(t *testing.T) {
	//通过
	var i64 int64 = 922337203685477583
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getLongFrom", context, i64)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt64(result))
	}
}
func TestGoiceGetFloat(t *testing.T) {
	//通过
	var f float32 = 234.43
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getFloat", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat32(result))
	}
}
func TestGoiceGetFloatUsingContext(t *testing.T) {
	//通过
	var f float32 = 23444444444.43
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getFloat", context, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat32(result))
	}
}
func TestGoiceGetFloatFrom(t *testing.T) {
	//通过
	var f float32 = 23444444444.43
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getFloatFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat32(result))
	}
}
func TestGoiceGetFloatFromUsingContext(t *testing.T) {
	//通过
	var f float32 = 23444444444.43
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getFloatFrom", context, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat32(result))
	}
}
func TestGoiceGetDouble(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getDouble", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}
func TestGoiceGetDoubleUsingContext(t *testing.T) {
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getDouble", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}
func TestGoiceGetDoubleFrom(t *testing.T) {
	//通过
	var f float64 = 43.43
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getDoubleFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}
func TestGoiceGetDoubleFromUsingContext(t *testing.T) {
	//通过
	var f float64 = 56.43
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "getDoubleFrom", context, f)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}
