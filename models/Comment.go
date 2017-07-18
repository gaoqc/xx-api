package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//文章模型

type Comment struct {
	TimeModel
	Author *User `orm:"rel(fk)"` //设置一对多关系
	//楼层
	Floor int64
	//评论内容
	Context string
	//评论所属文章
	Article *Article `orm:"rel(fk)"` //设置一对多关系
	//赞
	Good int64
}

func AddComment(v Comment) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&v)
	if err != nil {
		logs.Error("Insert Commenterr:%v", err)

	}
	return id
}
func CommentCount(articleId int) int64 {
	c, err := orm.NewOrm().QueryTable(Comment{}).Filter("article_id", articleId).Count()
	if err != nil {
		logs.Error("CommentCount err:%v", err)

	}
	return c
}

func ListComment(articleId int) []Comment {

	var list []Comment
	orm.NewOrm().QueryTable(Comment{}).Filter("article_id", articleId).All(&list)
	return list
}

func GetCommentById(id int) Comment {
	var c Comment
	orm.NewOrm().QueryTable(Comment{}).Filter("id", id).One(&c)
	return c
}

func DelComment(id, authorId int) int64 {
	num, err := orm.NewOrm().QueryTable(Comment{}).Filter("author_id", authorId).Filter("id", id).Delete()

	if err != nil {
		logs.Error("err at DelComment:%v ", err)
	}
	return num
}
