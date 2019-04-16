package main

import (
	"encoding/json"
	"fmt"
	"github.com/aloxc/goice/ice"
	"testing"
)

func TestGoiceExecute(t *testing.T) {
	//通过
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
	result := request.DoRequest(ice.ResponseType_Execute)
	ret := string(result)
	var response ice.Response
	json.Unmarshal(result, &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}
func TestGoiceExecuteUsingContext(t *testing.T) {
	//通过
	method := "getArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "123"

	context := make(map[string]string)
	context["name"] = "aloxc"

	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", context, req)
	result := request.DoRequest(ice.ResponseType_Execute)
	ret := string(result)
	var response ice.Response
	json.Unmarshal(result, &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}

func TestGoiceGetStringArticle(t *testing.T) {
	//通过
	method := "getStringArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "122"
	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
	result := request.DoRequest(ice.ResponseType_Execute)
	ret := string(result)
	var response ice.Response
	json.Unmarshal(result, &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}
func TestGoiceGetStringArticleUsingContext(t *testing.T) {
	//通过
	method := "getStringArticle"
	params := make(map[string]string)
	params["item"] = "free"
	params["id"] = "123"

	context := make(map[string]string)
	context["name"] = "aloxc"

	req := ice.NewReqeust(method, params)
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", context, req)
	result := request.DoRequest(ice.ResponseType_Execute)
	ret := string(result)
	var response ice.Response
	json.Unmarshal(result, &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}

func TestGoiceExecuteJson(t *testing.T) {
	//通过
	req := "{\"method\":\"getArticle\",\"params\":{\"item\":\"free\",\"id\":122222}}"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "executeJson", nil, req)
	result := request.DoRequest(ice.ResponseType_Execute_JSON)
	ret := string(result)
	var response ice.Response
	json.Unmarshal(result, &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}
func TestGoiceExecuteJsonUsingContext(t *testing.T) {
	//通过
	context := make(map[string]string)
	context["name"] = "aloxc"

	req := "{\"method\":\"getArticle\",\"params\":{\"item\":\"free\",\"id\":13333}}"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "executeJson", context, req)
	result := request.DoRequest(ice.ResponseType_Execute_JSON)
	ret := string(result)
	var response ice.Response
	json.Unmarshal(result, &response)
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}

func TestGoiceExecuteException(t *testing.T) {
	//通过
	req := "{\"method\":\"exception\",\"params\":{\"item\":\"free\",\"id\":122222}}"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "executeJson", nil, req)
	result := request.DoRequest(ice.ResponseType_Execute_JSON)
	ret := string(result)
	var response ice.Response
	err := json.Unmarshal(result, &response)
	if err != nil {
		fmt.Println("异常了", err)
	}
	fmt.Println(response)
	fmt.Println("请求结果", ret)
}
