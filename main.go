package main

import (
	"github.com/aloxc/goice/command"
	"github.com/aloxc/goice/ice"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

var ConnIce net.Conn
func init() {
	//日志格式为json
	log.SetFormatter(&log.TextFormatter{})
	//日志输出到标准控制台
	log.SetOutput(os.Stdout)
	//日志设置为warn级别以上
	log.SetLevel(log.TraceLevel)

}

func main() {
	log.Info("准备启动客户端，当前路径main" +os.Args[0])
	cli := command.Command{}
	cli.Run()
	time.Sleep(time.Second * 2)

	TestGoiceChinese2()
	//TestUserpostChinese2()
}
func TestUserpostChinese2() {
	log.Info("进入到Userpost测试代码中")
	start := time.Now().UnixNano()
	var times = 300000
	//测试通过aa

	for i := 0; i < times; i++ {
		request := ice.NewIceRequest("UserPostIce", ice.OperatorModeNormal, "threeparams", nil, "xyz", 4,float64(2))
		data,err := request.DoRequest(ice.ResponseType_String)
		//request.DoRequest(ice.ResponseType_String)
		if err != nil {
			log.Info(err)
		}
		if i%2000 == 0 {
			log.Info(i, (time.Now().UnixNano()-start)/1000000, "  ", data)
		}
	}
	log.Infof("执行Userpost[%d]，花费[%d]秒\n", times,(time.Now().UnixNano()-start)/1000000000)
}
func TestGoiceChinese2() {
	log.Info("进入到goice测试代码中")
	start := time.Now().UnixNano()
	var times = 3000000
	//测试通过aa
	for i := 0; i < times; i++ {
		request := ice.NewIceRequest("GoiceIce", ice.OperatorModeNormal, "two", nil, "我", "你")
		data,err := request.DoRequest(ice.ResponseType_String)
		//request.DoRequest(ice.ResponseType_String)
		if err != nil {
			log.Info(err)
		}
		if i%2000 == 0 {
			log.Info(i, (time.Now().UnixNano()-start)/1000000, "  ", data)
		}
	}
	log.Infof("执行goice[%d]，花费[%d]秒\n", times,(time.Now().UnixNano()-start)/1000000000)
}

