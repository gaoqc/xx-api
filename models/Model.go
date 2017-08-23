package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Model struct {
	Id int `orm:"pk;auto";json:"id"`
}
type TimeModel struct {
	Model
	Created time.Time `orm:"auto_now_add;type(datetime)";json:"created"`
	Updated time.Time `orm:"auto_now;type(datetime)";json:"updated"`
}

//新增记录
func AddOne(v interface{}) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(v)
	if err != nil {
		beego.Error("Insert err:%v", err)

	}
	return id

}

//批量新增
func AddMutil(v interface{}) {
	logs.Info("begin insert muilts")
	o := orm.NewOrm()
	_, err := o.InsertMulti(100, v)
	if err != nil {
		beego.Error("Insert muilts err:%v", err)

	}
}
