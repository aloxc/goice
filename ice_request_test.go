package main

import (
	"fmt"
	"github.com/aloxc/goice/ice"
	"github.com/aloxc/goice/utils"
	"reflect"
	"testing"
)

func TestGoiceSayHi(t *testing.T) {
	//goice.sayHi() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"sayHi",nil,"")
	result := request.DoRequest(ice.ResponseType_String)
	fmt.Println("请求结果",string(result))
}
func TestGoiceSayHiUsingContext(t *testing.T) {
	//goice.sayHi(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"sayHi",context,"")
	result := request.DoRequest(ice.ResponseType_String)
	fmt.Println("请求结果",string(result))
}
func TestGoiceSayHello(t *testing.T) {
	//hello.sayHello(string) 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"sayHello",nil,"aaa")
	result := request.DoRequest(ice.ResponseType_String)
	fmt.Println("请求结果",string(result))
}
func TestGoiceSayHelloUsingContext(t *testing.T) {
	//goice.sayHi(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"sayHello",context,"aaa")
	result := request.DoRequest(ice.ResponseType_String)
	fmt.Println("请求结果",string(result))
}

func TestGoiceVoid(t *testing.T) {
	//goice.vvoid() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"vvoid",nil,"")
	result := request.DoRequest(ice.ResponseType_Void)
	fmt.Println("请求结果",string(result))
}

func TestGoiceVoidUsingContext(t *testing.T) {
	//goice.vvoid(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"vvoid",context,"")
	result := request.DoRequest(ice.ResponseType_Void)
	fmt.Println("请求结果",string(result))
}

func TestGoiceVoidTo(t *testing.T) {
	//goice.vvoid(string) 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"vvoidTo",nil,"aaa")
	result := request.DoRequest(ice.ResponseType_Void)
	fmt.Println("请求结果",string(result))
}

func TestGoiceVoidToUsingContext(t *testing.T) {
	//goice.vvoid(string,context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"vvoidTo",context,"aaa")
	result := request.DoRequest(ice.ResponseType_Void)
	fmt.Println("请求结果",string(result))
}

func TestGoiceGetAge(t *testing.T) {
	//goice.getAge() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"getAge",nil,123321)
	result := request.DoRequest(ice.ResponseType_Int)
	fmt.Println("请求结果",utils.BytesToInt(result))
}

func TestGoiceGetAgeUsingContext(t *testing.T) {
	//goice.getAge(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"getAge",context,123)
	result := request.DoRequest(ice.ResponseType_Int)
	fmt.Println("请求结果",utils.BytesToInt(result))
}

func TestGoiceGetMaxAge(t *testing.T) {
	//goice.getAge() 测试成功
	var ii = []int{123,4,5}
	fmt.Println(reflect.TypeOf(ii))
	ix := []interface{}{1,2.3,"a"}
	fmt.Println(reflect.TypeOf(ix))

	//request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"getMaxAge",nil,123321)
	//result := request.DoRequest(ice.ResponseType_Int)
	//fmt.Println("请求结果",utils.BytesToInt(result))
}

func TestGoiceGetMaxAgeUsingContext(t *testing.T) {
	//goice.getAge(context) 测试成功
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"getMaxAge",context,123)
	result := request.DoRequest(ice.ResponseType_Int)
	fmt.Println("请求结果",utils.BytesToInt(result))
}
