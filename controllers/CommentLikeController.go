package controllers

/**

评论喜欢
**/
import (
	"github.com/astaxie/beego"
	"xx-api/models"
	"xx-api/utils"
)

type CommentLikeController struct {
	beego.Controller
}

/**
新增
**/
// @router /add [post]
func (c *CommentLikeController) Add() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	commentId, _ := c.GetInt("commentId")
	count := models.CountCommentLike(commentId)
	if models.AlreadyLike(commentId, user.Id) {

		c.Data["json"] = SuccessVO(count)

	} else {
		count = count + 1

		comment := models.GetCommentById(commentId)

		m := models.CommentLike{}
		m.LikeUser = &user
		m.Comment = &comment
		models.AddCommentLike(m)
	}

	c.Data["json"] = SuccessVO(count)
	c.ServeJSON()

}

/**
更新
**/
// @router /update [post]
func (c *CommentLikeController) Update() {
	var m models.CommentLike
	// user := GetUser(c.GetSession(utils.TicketName).(string))
	// m := models.Article{Author: &user, Title: c.GetString("title"), SubTitle: c.GetString("subTitle"), Keywords: c.GetString("keywords"), Context: c.GetString("context")}
	// logs.Debug("add article:%v", utils.ToJson(m))
	num := models.UpdateCommentLike(m)
	if num > 0 {
		c.Data["json"] = SuccessVO(nil)
	} else {
		c.Data["json"] = GetRetVO(UpdateFailCode, UpdateFailMsg, nil)
	}
	c.ServeJSON()

}

/**
查询
**/
// @router /list [get]
func (c *CommentLikeController) List() {
	list := models.ListCommentLike()

	c.Data["json"] = SuccessVO(list)
	c.ServeJSON()
}

/**

**/
// @router /del [post]
func (c *CommentLikeController) Del() {
	id, _ := c.GetInt("id")
	num := models.DelCommentLike(id)
	if num > 0 {
		c.Data["json"] = SuccessVO(nil)
	} else {
		c.Data["json"] = GetRetVO(UpdateFailCode, UpdateFailMsg, nil)
	}
	c.ServeJSON()
}
