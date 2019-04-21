package main

import (
	"encoding/json"
	"fmt"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/ice"
	"github.com/siddontang/go/log"
	"strconv"
	"testing"
	"time"
)

func init() {
	config.ReadConfig("")
}
func TestUserpostValuestr(t *testing.T) {

	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "valuestr", nil, "name", "aloxc")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestUserpostGetIntArray(t *testing.T) {
	//测试成功
	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "getIntArr", nil, 300)
	result, err := request.DoRequest(ice.ResponseType_Int_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestUserpostGetStrArray(t *testing.T) {
	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "getStrArr", nil, 300)
	result, err := request.DoRequest(ice.ResponseType_String_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceVoid(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "vvoid", nil, "")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceVoidTo(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "vvoidTo", nil, "aaa")
	result, err := request.DoRequest(ice.ResponseType_Void)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetString(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getStringFrom", nil, "ooye")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetStringArr(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getStringArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_String_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetStringArrFrom(t *testing.T) {
	//测试通过aa
	var arr = []string{"aaa", "bbb", "aaa"}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getStringArrFrom", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_String_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func BenchmarkGoicetwo(b *testing.B) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "two", nil, "aaa", "bbb")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoicetwo(t *testing.T) {

	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "two", nil, "aaa", "bbb")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceChineseq(t *testing.T) {
	start := time.Now().UnixNano()
	var times = 100000

	//测试通过aa
	for i := 0; i < times; i++ {
		request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "two", nil, "我"+strconv.Itoa(i), "你"+strconv.Itoa(i))
		_, err := request.DoRequest(ice.ResponseType_String)
		reError(err)
		//if showResult {
		//	log.Info("请求结果", result)
		//}
	}
	log.Infof("执行[%d]\n", times)
	log.Infof("flush=%d\n", time.Now().UnixNano()-start)
}
func TestGoiceGetBool(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBool", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetBoolFrom(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBoolFrom", nil, false)
	result, err := request.DoRequest(ice.ResponseType_Bool)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetBoolArray(t *testing.T) {
	//测试成功aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBoolArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Bool_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetBoolArrayFrom(t *testing.T) {
	//测试通过aa
	var arr = make([]bool, 10)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			arr[i] = true
		} else {
			arr[i] = false
		}
	}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getBoolArrFrom", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_Bool_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetByte(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByte", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetByteFrom(t *testing.T) {
	//测试通过aa
	var i8 int8 = 101
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByteFrom", nil, i8)
	result, err := request.DoRequest(ice.ResponseType_Int8)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetByteArray(t *testing.T) {
	//测试成功aa

	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByteArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int8_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetByteArrayFrom(t *testing.T) {
	//测试成功aa
	var arr = make([]int8, 10)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			arr[i] = 100 + int8(i)
		} else {
			arr[i] = -100 + int8(i)
		}
	}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getByteArrFrom", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_Int8_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetShort(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getShort", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetShortFrom(t *testing.T) {
	//测试通过aa
	var i16 int16 = 321
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getShortFrom", nil, i16)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetShortArray(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getShortArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int16_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetShortArrayFrom(t *testing.T) {
	//测试成功aa
	var arr = []int16{100, 1000, 10000}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getShortArrFrom", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_Int16_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetInt(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getInt", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetIntFrom(t *testing.T) {
	//测试通过aa
	var i16 int = 321
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getIntFrom", nil, i16)
	result, err := request.DoRequest(ice.ResponseType_Int16)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetIntArray(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getIntArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetIntArrayFrom(t *testing.T) {
	//测试成功aa
	var arr = []int{100, 1000, 10000}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getIntArrFrom", nil, arr)
	result, err := request.DoRequest(ice.ResponseType_Int_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetLong(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLong", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetLongFrom(t *testing.T) {
	//测试通过aa
	var i64 int64 = 922337203685477581
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLongFrom", nil, i64)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetArrLong(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLongArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Int64_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetLongArrFrom(t *testing.T) {
	//测试通过aa
	var i64 = []int64{922337203685477581, 5, 200000000}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLongArrFrom", nil, i64)
	result, err := request.DoRequest(ice.ResponseType_Int64_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}

func TestGoiceGetFloat32(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getFloat", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloat32From(t *testing.T) {
	//测试通过aa
	var f float32 = 444.43
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getFloatFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float32)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloat32Arr(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getFloatArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float32_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloatArr32From(t *testing.T) {
	//测试通过aa
	var f = []float32{33.43, 22, 3, 43}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getFloatArrFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float32_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloat64(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getDouble", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloat64From(t *testing.T) {
	//测试通过aa
	var f float64 = 4444.43
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getDoubleFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloat64Arr(t *testing.T) {
	//测试通过aa
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getDoubleArr", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
	}
}
func TestGoiceGetFloatArr64From(t *testing.T) {
	//测试通过aa
	var f = []float64{2.32, 44444.12, 344}
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getDoubleArrFrom", nil, f)
	result, err := request.DoRequest(ice.ResponseType_Float64_Array)
	reError(err)
	if showResult {
		log.Info("请求结果", result)
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
		log.Info("请求结果", result)
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

var showResult = false

func TestGoiceExecute(t *testing.T) {

	//通过
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	ret := result
	if showResult {
		var response ice.Response
		json.Unmarshal([]byte(result.(string)), &response)
		fmt.Println(response)
		log.Info("请求结果", ret)
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
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	ret := result
	if showResult {
		var response ice.Response
		json.Unmarshal([]byte(result.(string)), &response)

		fmt.Println(response)
		log.Info("请求结果", ret)
	}
}

func TestGoiceExecuteLargeString(t *testing.T) {
	//通过
	method := "getLargeString"
	params := make(map[string]string)
	params["a"] = "a"
	req := ice.NewReqeust(method, params)

	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	ret := result
	if showResult {
		log.Info("请求结果", ret)
	}
}
