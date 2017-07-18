package controllers

/**

评论相关
**/
import (
	"github.com/astaxie/beego"
	"xx-api/models"
	"xx-api/utils"
)

type CommentController struct {
	beego.Controller
}

/**
新增
**/
// @router /add [post]
func (c *CommentController) Add() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	articleId, _ := c.GetInt("articleId")
	article := models.GetArticleById(articleId)
	floor := models.CommentCount(articleId)

	model := models.Comment{Context: c.GetString("context")}

	model.Author = &user
	model.Article = &article
	model.Floor = floor

	id := models.AddComment(model)

	c.Data["json"] = SuccessVO(id)
	c.ServeJSON()

}

/**
查询
**/
// @router /list [get]
func (c *CommentController) List() {
	articleId, _ := c.GetInt("articleId")

	list := models.ListComment(articleId)

	c.Data["json"] = SuccessVO(list)
	c.ServeJSON()
}

/**

**/
// @router /del [post]
func (c *CommentController) Del() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	id, _ := c.GetInt("id")
	num := models.DelComment(id, user.Id)
	if num > 0 {
		c.Data["json"] = SuccessVO(nil)
	} else {
		c.Data["json"] = GetRetVO(UpdateFailCode, UpdateFailMsg, nil)
	}
	c.ServeJSON()
}
