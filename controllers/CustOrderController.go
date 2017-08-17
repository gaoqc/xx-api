package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"xx-api/models"
)

type CustOrderController struct {
	beego.Controller
}
type CustOrder struct {
	Name    string `json:"name";form:"name"`
	Phone   string `json:"phone";form:"phone"`
	Address string `json:"address";form:"address"`
	FixMsg  string `json:"fixMsg";form"fixMsg"`
	desc    string `json:"desc";form:"desc"`
}

// @router /addOrder [post]
func (c *CustOrderController) Add() {

	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	fixMsg := c.GetString("fixMsg")
	desc := c.GetString("desc")
	model := models.CustOrder{Name: name, Phone: phone, Address: address, FixMsg: fixMsg, Desc: desc}
	logs.Debug("cust order model:%v", model)
	models.AddCustOrder(model)
	c.Data["json"] = SuccessVO(CustOrder{name, phone, address, fixMsg, desc})

	// c.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.ServeJSON()

}
