package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 文章评论
type Comment struct {
	TimeModel
	author *User
	//楼层
	floor int64
	//评论内容
	context string
	//评论所属文章
	article *Article
}

func AddComment(comment Comment) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(comment)
	if err != nil {
		logs.Error("Insert err:")

	}
	return id
}
func ListComments(articleId int) []Comment {
	var list []Comment
	orm.NewOrm().QueryTable(Comment{}).Filter("article_id", articleId).All(&list)
	return list
}
func DelComment(commentId int) int64 {
	c := Comment{}
	c.Id = commentId
	o := orm.NewOrm()
	num, err := o.Delete(&c)
	if err != nil {
		logs.Error("err at del ")
	}
	return num
}
