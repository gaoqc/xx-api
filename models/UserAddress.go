package models

import (
	// "github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserAddress struct {
	TimeModel
	UserName      string `json:"userName"`
	UserPhone     string `json:"userPhone"`
	UserDomain    string `json:"userDomain"`
	UserAddDetail string `json:"userAddDetail"`
	User          *User  `orm:"rel(fk)"` //设置一对多关系
}

func ListAddress(userId int) []UserAddress {
	var list []UserAddress
	orm.NewOrm().QueryTable(UserAddress{}).Filter("user_id", userId).All(&list)
	return list
}
func QryAddr(addrId int) UserAddress {
	var addr UserAddress
	err := orm.NewOrm().QueryTable(UserAddress{}).Filter("id", addrId).One(&addr)
	if err != nil {
		logs.Error("qry err:%v", err)

	}
	return addr
}

//增加
func AddAddress(address *UserAddress) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(address)
	if err != nil {
		logs.Error("Insert err:%v", err)

	}
	return id
}

//更新
func UpdateAddress(address UserAddress) int64 {

	num, err := orm.NewOrm().Update(address)
	if err != nil {
		logs.Error("err at update user")
	}
	return num

}

//删除
func DelAddress(id int) int64 {
	add := UserAddress{}
	add.Id = id
	o := orm.NewOrm()
	num, err := o.Delete(&add)
	if err != nil {
		logs.Error("err at del userAddress")
	}
	return num
}
