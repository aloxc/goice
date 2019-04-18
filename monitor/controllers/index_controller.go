package controllers

import (
	"github.com/aloxc/goice/config"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Website"] = "goice"
	c.Data["appName"] = beego.BConfig.AppName
	c.Data["Email"] = "leerohwa@gmail.com"
	c.Data["config"] = config.ConfigMap
	c.TplName = "index.tpl"
}
