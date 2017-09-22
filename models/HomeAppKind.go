package models

//家电类型
import (
	// "github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	// "time"
)

type HomeAppKind struct {
	TimeModel
	Name string `json:"name"`
}

func QryHomeAppKind(id int) HomeAppKind {
	var v HomeAppKind
	orm.NewOrm().QueryTable(HomeAppKind{}).Filter("id", id).One(&v)
	return v
}
func QryAllHomeApp() []HomeAppKind {
	var vs []HomeAppKind
	orm.NewOrm().QueryTable(HomeAppKind{}).All(&vs)
	return vs
}
