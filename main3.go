package main

//发送40字节的head试试

import (
	"fmt"
	"github.com/aloxc/goice/ice"
)

func main3() {

	//hello.sayHello(string) 测试成功
	//request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"sayHello","aaa")
	//result := request.DoRequest()
	//fmt.Println("请求结果",string(result))

	//Goice.sayHi()无參及
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce",""),ice.OperatorModeNormal,"sayHi",nil,"")
	result := request.DoRequest(ice.ResponseType_String)
	fmt.Println("请求结果",string(result))
}

