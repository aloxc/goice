package command

import (
	"fmt"
	_ "github.com/aloxc/goice/banner"
	"github.com/aloxc/goice/config"
	"github.com/siddontang/go-log/log"
	"os"
	"time"
)

//解析参数

const (
	usage = `
参数列表
  -OperateTimeout int类型，ice请求超时时间，单位秒
  -ConnectTimeout int类型，ice连接超时时间，单位秒
  -MessageMaxSize int类型，ice请求最大网络传输量，单位字节，超过此设置会分多条发送数据
  -DefaultClientSize int类型，单个ice服务缺省连接数
  -WarnClientSize int类型，单个ice服务连接数从低于此设置到超过此设置后会输出一条日志；从超过超过该设置到低于此设置也会输出一条日志
  -MaxClientSize int类型，单个ice服务最大连接数
  -Compress boolean类型[true，false]，是否要启用压缩，设置此值会消耗部分cpu资源，但是可以节省带宽，超过1000字节该设置才会生效
  -RetryCount int类型，连接超时、执行超时重试次数
  -MonitorPort int类型，web界面显示统计的端口，默认不开启，只有此设置后才开启
  -Heartbeat int类型，心跳监测间隔时间，单位秒，默认3秒
  -Balance int类型，负载均衡，1：随机，2：轮询，3：哈希，
  -Report int类型，报警功能，0：不报警，1：邮件，2：短信，3：两者
  -ConfigFile string，配置文件路径
  可以在启动的时候带上这些参数，这些参数将是优先级最高的
  参数优先级说明：
	最高优先级：启动的时候命令行所带参数
	居中优先级：配置文件servers下面的参数
	最低优先级：配置文件中default下面的参数
`
)

type Command struct {
}

func (this *Command) Run() {
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-help" || os.Args[1] == "help") {
		fmt.Println(usage)
		os.Exit(1)
	}
	var configFile = ""
	if len(os.Args) > 1 && (os.Args[1] == "--ConfigFile" || os.Args[1] == "-ConfigFile" || os.Args[1] == "ConfigFile") {
		configFile = os.Args[2]
	}
	config.ReadConfig(configFile)

	if config.MonitorPorta > 0 {
		startMonitor()
	}
	if config.ReportType > 0 {
		startReport()
	}
}

//报警功能
func startReport() {
	go func() {
		for {
			select {
			case <-time.After(time.Second * 3):
				log.Info("开始收集数据")

			}
		}
	}()
}

//启动http服务
func startMonitor() {

}
