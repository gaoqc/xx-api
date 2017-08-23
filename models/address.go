package models

import (
	"github.com/astaxie/beego/orm"
)

//省份
type Province struct {
	TimeModel
	Code string `json:"code"`
	Name string `json:"name"`
}

//地市
type City struct {
	Province
	ParentCode string `json:"parent_code"`
}

//区域
type Area struct {
	City
}

//街道
type Street struct {
	City
}

//查询所有省份
func QryAllProvince() []Province {
	var list []Province
	orm.NewOrm().QueryTable(Province{}).All(&list)
	return list
}

//查询省份下所有地市
func QryCityByProId(provinceId int) []City {
	var list []City
	orm.NewOrm().QueryTable(City{}).Filter("parent_code", provinceId).All(&list)
	return list
}

//查询地市下的所有区域
func QryAreaByCityId(cityId int) []Area {
	var list []Area
	orm.NewOrm().QueryTable(Area{}).Filter("parent_code", cityId).All(&list)
	return list

}

//查询区域下的所有街道
func QryStreetByArea(areaId int) []Street {
	var list []Street
	orm.NewOrm().QueryTable(Street{}).Filter("parent_code", areaId).All(&list)
	return list
}
