package main

import (
	"github.com/aloxc/goice/command"
	"github.com/aloxc/goice/ice"
	"github.com/siddontang/go-log/log"
	"strconv"
	"time"
)

func main() {
	cli := command.Command{}
	cli.Run()
	TestGoiceChinese()
}
func TestGoiceChinese() {
	start := time.Now().UnixNano()
	var times = 100000

	//测试通过aa
	for i := 0; i < times; i++ {
		request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "two", nil, "我"+strconv.Itoa(i), "你"+strconv.Itoa(i))
		_, err := request.DoRequest(ice.ResponseType_String)
		if err != nil {
			log.Info(err)
		}
		if i%2000 == 0 {
			log.Info(i, (time.Now().UnixNano()-start)/1000000)
		}
		//if showResult {
		//	log.Info("请求结果", result)
		//}
	}
	log.Infof("执行[%d]\n", times)
	log.Infof("flush=%d\n", time.Now().UnixNano()-start)
}
