package main

import (
	"github.com/aloxc/goice/command"
	"github.com/aloxc/goice/config"
	"github.com/aloxc/goice/ice"
	"github.com/siddontang/go-log/log"
	"net"
	"time"
)

var ConnIce net.Conn

func main() {
	log.Info("准备允许客户端")
	cli := command.Command{}
	cli.Run()
	TestGoiceChinese2()
}
func TestGoiceChinese2() {
	log.Info("进入到测试代码中")
	start := time.Now().UnixNano()
	var times = 100000
	//测试通过aa
	for i := 0; i < times; i++ {
		request := ice.NewIceRequest("Goice", ice.OperatorModeNormal, "two", nil, "我", "你")
		//data,err := request.DoRequest(ice.ResponseType_String)
		request.DoRequest(ice.ResponseType_String)
		//if err != nil {
		//	log.Info(err)
		//}
		//if i%2000 == 0 {
		//	log.Info(i, (time.Now().UnixNano()-start)/1000000, "  ", data)
		//}
		//if showResult {
		//	log.Info("请求结果", result)
		//}
	}
	log.Infof("执行[%d]\n", times)
	log.Infof("flush=%d\n", (time.Now().UnixNano()-start)/1000000000)
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
