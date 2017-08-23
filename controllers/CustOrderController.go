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
	AppTypeId int
	VendorId  int
	FixMsg    string
	AddrId    int
}

// @router /addOrder [post]
func (c *CustOrderController) Add() {

	appTypeId, _ := c.GetInt("appTypeId")
	vendorId, _ := c.GetInt("vendorId")
	fixMsg := c.GetString("fixMsg")
	addrId, _ := c.GetInt("addrId")

	model := models.CustOrder{AppTypeId: appTypeId, VendorId: vendorId, FixMsg: fixMsg, AddrId: addrId}
	logs.Info("order:%v", model)

	models.AddOne(&model)
	c.Data["json"] = SuccessVO("")

	// logs.Debug("cust order model:%v", model)
	// models.AddCustOrder(model)
	// c.Data["json"] = SuccessVO(CustOrder{name, phone, address, fixMsg, desc})

	// // c.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.ServeJSON()

}
