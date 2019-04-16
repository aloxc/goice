package main

import (
	"encoding/json"
	"fmt"
	"github.com/aloxc/goice/ice"
	"testing"
)

func reError(err error) {
	if err != nil {
		switch err.(type) {
		case *ice.UserError:
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
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
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
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
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

func TestGoiceExecuteJson(t *testing.T) {
	//通过
	req := "{\"method\":\"getArticle\",\"params\":{\"item\":\"free\",\"id\":122222}}"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "executeJson", nil, req)
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
func TestGoiceExecuteJsonUsingContext(t *testing.T) {
	//通过
	context := make(map[string]string)
	context["name"] = "aloxc"

	req := "{\"method\":\"getArticle\",\"params\":{\"item\":\"free\",\"id\":13333}}"
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "executeJson", context, req)
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

	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "execute", nil, req)
	result, err := request.DoRequest(ice.ResponseType_Execute_JSON)
	reError(err)
	ret := string(result)
	if showResult {
		fmt.Println("请求结果", ret)
	}
}
