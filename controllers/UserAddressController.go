package controllers

import (
	"github.com/astaxie/beego"

	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"xx-api/models"
	"xx-api/utils"
)

type UserAddressController struct {
	beego.Controller
}
type AddressForm struct {
	UserName      string `form:"userName";valid:"Required"`
	UserPhone     string `form:"userPhone";valid:"Mobile"`
	UserDomain    string `form:"userDomain";valid:"Required"`
	UserAddDetail string `form:"userAddDetail";valid:"Required"`
}

//入参验证
func (c *AddressForm) Valid(v *validation.Validation) {

}

/**
**新增
**/
// @router  /add [post]
func (c *UserAddressController) AddAddress() {
	var form AddressForm
	c.ParseForm(&form)
	msg := ValidParams(&form)
	if len(msg) > 0 {
		c.Data["json"] = msg
		c.ServeJSON()
		return
	}
	ticket := c.GetSession(utils.TicketName)
	user := GetUser(ticket.(string))

	address := models.UserAddress{UserName: form.UserName, UserPhone: form.UserPhone, UserDomain: form.UserDomain, UserAddDetail: form.UserAddDetail}
	address.User = &user

	id := models.AddAddress(&address)
	c.Data["json"] = SuccessVO(id)
	c.ServeJSON()
}

// @router /list [get]
func (c *UserAddressController) ListAddress() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	list := models.ListAddress(user.Id)

	c.Data["json"] = SuccessVO(list)
	c.ServeJSON()
}

// @router /del [get]
func (c *UserAddressController) DelAddress() {
	id, _ := c.GetInt("id")
	num := models.DelAddress(id)
	if num > 0 {
		c.Data["json"] = SuccessVO(nil)
	} else {
		c.Data["json"] = GetRetVO(UpdateFailCode, UpdateFailMsg, nil)
	}
	c.ServeJSON()
}
