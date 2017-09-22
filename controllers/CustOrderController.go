package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
	"xx-api/models"
	"xx-api/utils"
)

type CustOrderController struct {
	beego.Controller
}
type CustOrder struct {
	AppTypeId int
	VendorId  int
	FixMsg    string
	AddrId    int
	UserId    int
}

// @router /addOrder [post]
func (c *CustOrderController) Add() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	kindId, _ := c.GetInt("appTypeId")
	homekind := models.QryHomeAppKind(kindId)
	vendorId, _ := c.GetInt("vendorId")
	vendor := models.QryVendor(vendorId)
	addrId, _ := c.GetInt("addrId")
	addr := models.QryAddr(addrId)
	fixMsg := c.GetString("fixMsg")
	model := models.CustOrder{User: &user, AppKind: &homekind, Vendor: &vendor, FixMsg: fixMsg, Addr: &addr}
	logs.Info("order:%v", model)

	models.AddOne(&model)
	c.Data["json"] = SuccessVO("")

	// logs.Debug("cust order model:%v", model)
	// models.AddCustOrder(model)
	// c.Data["json"] = SuccessVO(CustOrder{name, phone, address, fixMsg, desc})

	// // c.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.ServeJSON()

}

// @router /qryOrderList [get]
func (c *CustOrderController) QryOrderList() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	ret := models.QryCustOrderByUserId(user.Id)
	c.Data["json"] = SuccessVO(toVO(ret))
	c.ServeJSON()
}

type OrderVo struct {
	AppName    string    `json:"appName"`
	VendorName string    `json:"vendorName"`
	CreateTime time.Time `json:"createTime"`
	Status     string    `json:"status"`
	FixMsg     string    `json:"fixMsg"`
}

//状态:0,待处理,1正在处理,2处理完成
func getStatusCn(status int) string {
	cn := "待处理"
	switch status {
	case 1:
		cn = "处理中"
	case 2:
		cn = "已完成"
	default:
	}
	return cn
}

func toVO(orders []models.CustOrder) []OrderVo {

	list := make([]OrderVo, len(orders))
	for index, order := range orders {
		logs.Info("to vo:%v", order.Id)
		// logs.Info(utils.ToJson(order))
		vo := OrderVo{AppName: order.AppKind.Name, VendorName: order.Vendor.Name, CreateTime: order.Created, Status: getStatusCn(order.Status)}
		vo.FixMsg = order.FixMsg
		list[index] = vo
	}
	logs.Info("len(list):%v", len(list))
	return list

}
