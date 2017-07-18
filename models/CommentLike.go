
package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//文章模型

type CommentLike struct {
	TimeModel
	 
}

func AddCommentLike(v CommentLike) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&v)
	if err != nil {
		logs.Error("Insert CommentLikeerr:%v",err)

	}
	return id
}
func UpdateCommentLike(v CommentLike) int64{
	up := make(map[string]interface{})
	 
	num, err := orm.NewOrm().QueryTable(CommentLike{}).Filter("id", v.Id).Update(up)
	if err != nil {
		logs.Error("err at UpdateCommentLike:%v",err)
	}
	return num
}
func ListCommentLike()[]CommentLike {
	
	var list []CommentLike
    orm.NewOrm().QueryTable(CommentLike{}).Filter("id", "").All(&list)
    return list
}
func DelCommentLike(id int)int64 {
	  num, err := orm.NewOrm().QueryTable(CommentLike{}).Filter("id", id).Delete()
      if err != nil {
      	logs.Error("err at DelCommentLike:%v ", err)
       }
      return num
}

