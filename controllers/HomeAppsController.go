package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
)

type HomeAppsController struct {
	beego.Controller
}

type HomeApp struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Src  string `json:"src"`
}

// @router /qryAllHomeAppliances [get]
func (c *HomeAppsController) QryAllHomeAppliances() {
	// types := c.GetString("type")
	bx := HomeApp{"冰箱", "bx", "bingxiang.png"}
	kt := HomeApp{"空调", "kt", "kongtiao.png"}
	xyj := HomeApp{"洗衣机", "xyj", "xiyiji.png"}
	dsj := HomeApp{"电视机", "dsj", "dsj.png"}
	rsq := HomeApp{"热水器", "rsq", "rsq.png"}
	apps := []HomeApp{bx, kt, xyj, dsj, rsq}
	c.Data["json"] = SuccessVO(apps)
	c.ServeJSON()

}
