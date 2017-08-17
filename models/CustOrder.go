package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// "time"
)

//文章模型

type CustOrder struct {
	TimeModel
	Name    string
	Phone   string
	Address string
	Status  int `orm:"default(0)"` //状态:0,未处理,1正在处理,2处理完成
	FixMsg  string
	// HaveTime time.Time
	Desc string
}

func AddCustOrder(v CustOrder) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&v)
	if err != nil {
		logs.Error("Insert CustOrder err:%v", err)

	}
	return id
}
func qryCustOrderByPhone(phone string) []CustOrder {
	var list []CustOrder
	orm.NewOrm().QueryTable(CustOrder{}).Filter("phone", phone).All(&list)
	return list
}
