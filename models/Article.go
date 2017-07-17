package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//文章模型

type Article struct {
	TimeModel
	//标题
	Title string
	//副标题
	SubTitle string
	//内容
	Context string
	//所属人
	Author *User `orm:"rel(fk)"` //设置一对多关系
	//文章关键字
	Keywords string
}

func AddArticle(v Article) int64 {

	o := orm.NewOrm()
	id, err := o.Insert(&v)
	if err != nil {
		logs.Error("Insert Articleerr:%v", err)

	}
	return id
}
func ListArticle(authorId int) []Article {

	var list []Article
	orm.NewOrm().QueryTable(Article{}).Filter("author_id", authorId).All(&list)
	return list
}
func DelArticle(id int, authorId int) int64 {
	c := Article{}
	c.Id = id

	o := orm.NewOrm()
	num, err := o.QueryTable(c).Filter("author_id", authorId).Filter("id", id).Delete()

	if err != nil {
		logs.Error("err at del ")
	}
	return num
}
