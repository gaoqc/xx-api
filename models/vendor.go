package models

import (
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// "time"
)

type Vendor struct {
	TimeModel
	Name string `json:"name"`
}

func QryVendor(id int) Vendor {
	var v Vendor
	orm.NewOrm().QueryTable(Vendor{}).Filter("id", id).One(&v)
	return v
}
func QryAllVendors() []Vendor {
	var vs []Vendor
	orm.NewOrm().QueryTable(Vendor{}).All(&vs)
	return vs
}
