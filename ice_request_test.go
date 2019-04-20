package main

import (
	"encoding/json"
	"fmt"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"testing"
)

func init() {
	config.ReadConfig("")
}
func TestUserpostValuestr(t *testing.T) {
	//goice.sayHi() 测试成功
	//params := []string{"a", "bc"}

	params := &[]interface{}{"a", "bc", 1}
	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "valuestr", nil, params)
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}

func TestUserpostGetIntArr(t *testing.T) {
	//goice.sayHi() 测试成功
	//params := []string{"a", "bc"}

	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "getIntArr", nil, 300)
	result, err := request.DoRequest(ice.ResponseType_Int_Array)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestUserpostGetStrArr(t *testing.T) {
	//goice.sayHi() 测试成功
	//params := []string{"a", "bc"}

	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "getStrArr", nil, 300)
	result, err := request.DoRequest(ice.ResponseType_String_Array)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}

func TestGoiceVoid(t *testing.T) {
	//goice.vvoid() 测试成功
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "vvoid", nil, "")
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
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "vvoid", context, "")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceVoidTo(t *testing.T) {
	//goice.vvoid(string) 测试成功
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "vvoidTo", nil, "aaa")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
func TestGoiceGetBool(t *testing.T) {
	//通过
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBool", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetBoolFrom(t *testing.T) {
	//通过
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBoolFrom", nil, false)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetBoolArr(t *testing.T) {
	//通过
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBoolArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Bool_Array)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetBoolArrFrom(t *testing.T) {
	//通过
	var arr = make([]bool, 10)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			arr[i] = true
		} else {
			arr[i] = false
		}
	}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBoolArr", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_Bool_Array)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToBool(result))
	}
}
func TestGoiceGetByte(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功
	var i8 int8 = 101
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByte", nil, i8)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt8(result))
	}
}
func TestGoiceGetByteFrom(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功
	var i8 int8 = 101
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByte", nil, i8)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt8(result))
	}
}
func TestGoiceGetByteArr(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功

	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByteArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt8(result))
	}
}
func TestGoiceGetByteArrFrom(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功
	var arr = make([]int8, 10)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			arr[i] = 100 + int8(i)
		} else {
			arr[i] = -100 + int8(i)
		}
	}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByteArrFrom", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt8(result))
	}
}
func TestGoiceGetShort(t *testing.T) {
	//goice.getShortFrom(int16) 测试成功
	//var i16 int16 = 321
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getShort", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt16(result))
	}
}
func TestGoiceGetShortFrom(t *testing.T) {
	//goice.getShortFrom(context) 测试成功
	var i16 int16 = 321
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getShortFrom", nil, i16)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt16(result))
	}
}
func TestGoiceGetLong(t *testing.T) {
	//通过
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLong", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt64(result))
	}
}
func TestGoiceGetLongFrom(t *testing.T) {
	//通过
	var i64 int64 = 922337203685477581
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLongFrom", nil, i64)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.BytesToInt64(result))
	}
}
func TestGoiceGetFloat(t *testing.T) {
	//通过
	var f float32 = 234.43
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getFloat", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat32(result))
	}
}
func TestGoiceGetFloatFrom(t *testing.T) {
	//通过
	var f float32 = 23444444444.43
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getFloatFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat32(result))
	}
}
func TestGoiceGetDouble(t *testing.T) {
	//通过
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getDouble", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}
func TestGoiceGetDoubleFrom(t *testing.T) {
	//通过
	var f float64 = 43.43
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getDoubleFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}

// ice 服务器端 业务逻辑异常
func TestGoiceExecuteException(t *testing.T) {
	//通过
	method := "exception"
	params := make(map[string]string)
	params["type"] = "zero"
	//params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	_, err := request.DoRequest(ice.ResponseType_Execute)
	//reError(err)
	fmt.Println(err)
}

// ice 超时
func TestGoiceExecuteTimeout(t *testing.T) {
	//通过
	method := "exception"
	params := make(map[string]string)
	params["type"] = "timeout"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	_, err := request.DoRequest(ice.ResponseType_Execute)
	//reError(err)
	fmt.Println(err)
}

//调用方法不存在异常
func TestGoiceOperatoNotExist(t *testing.T) {
	//通过
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "abcdef", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	//reError(err)
	fmt.Println(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}

func reError(err error) {
	if err != nil {
		switch err.(type) {
		case *ice.UserUnknownError:
			break
		default:
			panic(err)
		}
	}
}

var showResult = true

func TestGoiceExecute(t *testing.T) {

	//通过
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_Execute)
	reError(err)
	ret := string(result)
	if showResult {
		var response ice.Response
		json.Unmarshal(result, &response)
		fmt.Println(response)
		fmt.Println("请求结果", ret)
	}
}

func TestGoiceGetStringArticle(t *testing.T) {
	//通过
	method := "getStringArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_Execute)
	reError(err)
	ret := string(result)
	if showResult {
		var response ice.Response
		json.Unmarshal(result, &response)
		fmt.Println(response)
		fmt.Println("请求结果", ret)
	}
}

func TestGoiceExecuteLargeString(t *testing.T) {
	//通过
	method := "getLargeString"
	params := make(map[string]string)
	params["a"] = "a"
	req := ice.NewReqeust(method, params)

	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_Execute_JSON)
	reError(err)
	ret := string(result)
	if showResult {
		fmt.Println("请求结果", ret)
	}
}
