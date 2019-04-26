package main

import (
	"github.com/aloxc/goice/command"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/ice"
	"github.com/siddontang/go-log/log"
	"net"
	"os"
	"time"
)

var ConnIce net.Conn

func main() {
	log.Info("准备启动客户端，当前路径main" +os.Args[0])
	cli := command.Command{}
	cli.Run()
	time.Sleep(time.Second * 2)

	TestGoiceChinese2()
	TestUserpostChinese2()
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
	var times = 300000
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

func TestGoiceChinese() {
	start := time.Now().UnixNano()
	var times = 1000000
	var err error
	ConnIce, err = net.DialTimeout("tcp", config.ConfigMap["Goice"][config.Address].(string), time.Second*5)
	if err != nil { //如果连接失败。则返回。
		log.Error(err)
		return
	}
	defer func() {
		time.Sleep(time.Second * 10)
		//if count % 1000 ==0{
		log.Info("关闭连接")
		ConnIce.Close()
		time.Sleep(time.Second * 20)
		//conn = nil
		//}
		//curPool.Return(conn)
	}()
	//测试通过aa
	for i := 0; i < times; i++ {
		request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "two", nil, "我", "你")
		data, err := request.DoRequest(ice.ResponseType_String)
		if err != nil {
			log.Info(err)
		}
		if i%2000 == 0 {
			log.Info(i, (time.Now().UnixNano()-start)/1000000, "  ", data)
		}
		//if showResult {
		//	log.Info("请求结果", result)
		//}
	}
	log.Infof("执行[%d]\n", times)
	log.Infof("flush=%d\n", (time.Now().UnixNano()-start)/1000000000)
}
