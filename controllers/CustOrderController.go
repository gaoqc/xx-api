package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CustOrderController struct {
	beego.Controller
}
type CustOrder struct {
	Name    string
	Phone   string
	Address string
}

// @router /add [GET]
func (c *CustOrderController) Add() {
	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	var cust CustOrder
	cust.Name = name
	cust.Phone = phone
	cust.Address = address
	// cust := CustOrder({name:name,phone:phone,address:address})
	logs.Info("name:%s,phone:%s,address:%s", name, phone, address)
	logs.Info("cust:%+v", cust)
	c.Data["json"] = `{name:'gaoqc', phone:'17704027838'}`

	// c.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.ServeJSON()

}
