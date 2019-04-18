package routers

import (
	"github.com/aloxc/goice/monitor/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/monitor", &controllers.MonitorController{}, "get:GetMonitor")
}
