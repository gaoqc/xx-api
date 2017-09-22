package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"xx-api/models"
)

type VendorController struct {
	beego.Controller
}

//@router /qryVendor [get]
func (c *VendorController) QryVendor() {
	id, _ := c.GetInt("vendorId")
	if 0 == id {
		c.Data["json"] = SuccessVO(models.QryAllVendors())
	} else {
		c.Data["json"] = SuccessVO(models.QryVendor(id))
	}
	c.ServeJSON()

}
