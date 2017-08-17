package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// "xx-api/models"
	// "xx-api/utils"
)

type IndexController struct {
	beego.Controller
}

// @router / [get]
func (c *IndexController) Index() {
	logs.Info("enter index")
	c.TplName = "index.html"
	c.TplExt = "html"
	c.Render()
}
