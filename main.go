package main

import (
	"encoding/json"
	"fmt"
	"github.com/aloxc/goice/command"
	"github.com/aloxc/goice/ice"
	"github.com/siddontang/go-log/log"
	"time"
)
import _ "github.com/aloxc/goice/ice"

func main() {
	cmd := command.Command{}
	cmd.Run()

	TestGoiceSayHi1()
	TestGoiceGetLongUsingContext1()
	TestGoiceExecute1()
	TestUserpostTodo()
	time.Sleep(time.Second * 3)
}

func reError3(err error) {
	if err != nil {
		switch err.(type) {
		case *ice.UserUnknownError:
			break
		default:
			panic(err)
		}
	}
}
func TestUserpostTodo() {
	//goice.sayHi() 测试成功
	request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "todo", nil, "this is a json?")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError3(err)
	//if showResult {
	//	if showResult {
	fmt.Println("请求结果", result)
	//}
	//}
}
func TestGoiceSayHi1() {
	//goice.sayHi() 测试成功
	log.Info("准备调用")
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "sayHi", nil, "")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError3(err)
	//if showResult {
	//	if showResult {
	fmt.Println("请求结果", result)
	//}
	//}
}
func TestGoiceGetLongUsingContext1() {
	//通过
	context := make(map[string]string)
	context["name"] = "aloxc"
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "getLong", context, nil)
	result, err := request.DoRequest(ice.ResponseType_Int64)
	reError3(err)
	fmt.Println("请求结果", result)

}
func TestGoiceExecute1() {

	//通过
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_Execute)
	reError3(err)
	ret := result
	var response ice.Response
	json.Unmarshal([]byte(result.(string)), &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}
