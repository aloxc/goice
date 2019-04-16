package main

//所有异常测试都在本文件中
import (
	"fmt"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"testing"
)

// ice 服务器端 业务逻辑异常
func TestGoiceExecuteException(t *testing.T) {
	//通过
	method := "exception"
	params := make(map[string]string)
	params["type"] = "zero"
	//params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
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
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
	_, err := request.DoRequest(ice.ResponseType_Execute)
	//reError(err)
	fmt.Println(err)
}

//调用方法不存在异常
func TestGoiceOperatoNotExist(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "abcdef", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_Float64)
	//reError(err)
	fmt.Println(err)
	if showResult {
		fmt.Println("请求结果", utils.ByteToFloat64(result))
	}
}

//ice协议异常
func TestGoiceFactNotExists(t *testing.T) {
	//通过
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "sayHi", nil, nil)
	result, err := request.DoRequest(ice.ResponseType_String)
	//reError(err)
	fmt.Println(err)
	if showResult {
		fmt.Println("请求结果", string(result))
	}
}
