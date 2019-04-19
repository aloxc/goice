package main

import (
	"fmt"
	"github.com/aloxc/goice/command"
	"github.com/aloxc/goice/ice"
	"time"
)
import _ "github.com/aloxc/goice/ice"

func main() {
	cmd := command.Command{}
	cmd.Run()

	TestGoiceSayHi()
	time.Sleep(time.Second * 3)
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
func TestGoiceSayHi() {
	//goice.sayHi() 测试成功
	request := ice.NewIceRequest(ice.NewIdentity("HelloIce", ""), ice.OperatorModeNormal, "sayHi", nil, "")
	result, err := request.DoRequest(ice.ResponseType_String)
	reError(err)
	//if showResult {
	//	if showResult {
	fmt.Println("请求结果", string(result))
	//}
	//}
}
