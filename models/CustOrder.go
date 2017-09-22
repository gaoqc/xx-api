package models

import (
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// "time"
)

//文章模型

type CustOrder struct {
	TimeModel
	AppKind *HomeAppKind `orm:"rel(fk)"`
	Vendor  *Vendor      `orm:"rel(fk)"`
	FixMsg  string
	Addr    *UserAddress `orm:"rel(fk)"`
	Status  int          `orm:"default(0)"` //状态:0,未处理,1正在处理,2处理完成
	User    *User        `orm:"rel(fk)"`
	Desc    string
}

// func AddCustOrder(v CustOrder) int64 {
// 	o := orm.NewOrm()
// 	id, err := o.Insert(&v)
// 	if err != nil {
// 		logs.Error("Insert CustOrder err:%v", err)

// 	}
// 	return id
// }
func qryCustOrderByPhone(phone string) []CustOrder {
	var list []CustOrder
	orm.NewOrm().QueryTable(CustOrder{}).Filter("phone", phone).All(&list)
	return list
}
func QryCustOrderByUserId(userId int) []CustOrder {
	var list []CustOrder
	orm.NewOrm().QueryTable(CustOrder{}).Filter("user_id", userId).RelatedSel().All(&list)
	// for _, order := range list {
	// 	logs.Info("order.vendor.name:%s", order.Vendor.Name)
	// 	logs.Info("order.vendor.Id:%v", order.Vendor.Id)
	// }
	return list
}
