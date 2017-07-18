package controllers

/**

文章相关
**/
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"xx-api/models"
	"xx-api/utils"
)

type ArticleController struct {
	beego.Controller
}
type ArticleForm struct {
	Title string `form:"title"`
	//副标题
	SubTitle string `form:"subTitle"`
	//内容
	Context string `form:"context"`

	//文章关键字
	Keywords string `form:"keywords"`
}

/**
新增
**/
// @router /add [post]
func (c *ArticleController) Add() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	var form ArticleForm
	c.ParseForm(&form)
	logs.Debug("artcle form:%v", utils.ToJson(form))
	// model := models.Article{Author: &user, Title: c.GetString("title"), SubTitle: c.GetString("subTime"), Keywords: c.GetString("keywords"), Context: c.GetString("context")}
	model := models.Article{Author: &user, Title: form.Title, SubTitle: form.SubTitle, Keywords: form.Keywords, Context: form.Context}
	logs.Debug("add article:%v", utils.ToJson(model))
	id := models.AddArticle(model)

	c.Data["json"] = SuccessVO(id)
	c.ServeJSON()

}

/**
更新
**/
// @router /update [post]
func (c *ArticleController) Update() {

	user := GetUser(c.GetSession(utils.TicketName).(string))
	m := models.Article{Author: &user, Title: c.GetString("title"), SubTitle: c.GetString("subTitle"), Keywords: c.GetString("keywords"), Context: c.GetString("context")}
	logs.Debug("add article:%v", utils.ToJson(m))
	num := models.UpdateArticle(m)
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
func (c *ArticleController) List() {
	authorId, _ := c.GetInt("authorId")
	list := models.ListArticle(authorId)

	c.Data["json"] = SuccessVO(list)
	c.ServeJSON()
}

/**
查询
**/
// @router /myList [get]
func (c *ArticleController) Myist() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	list := models.ListArticle(user.Id)

	c.Data["json"] = SuccessVO(list)
	c.ServeJSON()
}

/**

**/
// @router /del [post]
func (c *ArticleController) Del() {
	user := GetUser(c.GetSession(utils.TicketName).(string))
	id, _ := c.GetInt("id")
	num := models.DelArticle(id, user.Id)
	if num > 0 {
		c.Data["json"] = SuccessVO(nil)
	} else {
		c.Data["json"] = GetRetVO(UpdateFailCode, UpdateFailMsg, nil)
	}
	c.ServeJSON()
}
