package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	// "github.com/garyburd/redigo/redis"
	// "fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
	"xx-api/models"
	"xx-api/utils"
)

type UserForm struct {
	LoginAcc string `form:"loginAcc" ;valid:"Required"`
	LoginPwd string `form:"loginPwd";valid:"Required"`
	Phone    string `form:"phone";valid:"Mobile"`
	Email    string `form:"email";valid:"Email; MaxSize(100)"`
	TrueName string `form:trueName`
	IdNo     string `form:idNo`
	IdType   int    `form:idType`
	//用户类型,0普通顾客,1表示维修人员
	UserType int `form:"userType"`
}

type UserController struct {
	beego.Controller
}

//入参验证
func (c *UserForm) Valid(v *validation.Validation) {

}

// @router  /update [get]
func (c *UserController) Update() {
	idType, _ := c.GetInt("idType")
	userType, _ := c.GetInt("uerType")
	num := models.Update(c.GetString("phone"), c.GetString("email"), c.GetString("idNo"), c.GetString("trueName"), c.GetString("loginAcc"), idType, userType)
	if num > 0 {
		c.Ctx.WriteString("update ok!")
	} else {
		c.Ctx.WriteString("update fail,noChange")
	}

}

//注册新用户
// @router  /register [post]
func (c *UserController) Register() {

	//校验入参
	var userForm UserForm
	c.ParseForm(&userForm)

	valid := validation.Validation{}
	b, _ := valid.Valid(&userForm)
	var msg = ""
	if !b {
		// validation does not passRegistExistsCode
		// blabla...
		for _, err := range valid.Errors {
			msg = msg + err.Key + ":" + err.Message + "\n\r"
		}
		c.Ctx.WriteString(msg)
		return
	}
	//判断是否存在
	exists := models.ExistUser(userForm.LoginAcc, userForm.IdNo, userForm.Phone, userForm.Email)
	if exists {
		c.Data["json"] = GetRetVO(RegistExistsCode, RegistExistsMsg, nil)
		c.ServeJSON()
		return
	}

	//新增记录
	id := models.AddUser(&models.User{LoginAcc: userForm.LoginAcc, LoginPwd: userForm.LoginPwd, Phone: userForm.Phone, Email: userForm.Email, TrueName: userForm.TrueName, IdNo: userForm.IdNo, IdType: userForm.IdType, UserType: userForm.UserType})
	c.Data["json"] = SuccessVO("id =" + strconv.FormatInt(id, 10))
	c.ServeJSON()

}

// @router /login [get]
func (c *UserController) Login() {

	//如果是重新登陆,则直接返回
	utils.RedisDo("del", c.GetSession(utils.TicketName))

	c.DelSession(utils.TicketName)

	loginAcc := c.GetString("loginAcc")
	loginPwd := c.GetString("loginPwd")
	logs.Info("login with acc:%s", loginAcc)
	exist, user := models.Login(loginAcc, loginPwd)
	if exist {
		ticket := utils.CreateTicket()
		c.SetSession(utils.TicketName, ticket)
		c.Data["json"] = SuccessVO(user)
		// utils.RedisDo("set", ticket, user, "EX", strconv.FormatInt(60*30, 10))
		utils.RedisSet(ticket, user, "EX", strconv.FormatInt(60*30, 10))

	} else {
		c.Data["json"] = GetRetVO(LoginFailCode, LoginFailMsg, nil)
	}
	c.ServeJSON()

}

//退出登陆
// @router /logout [post]
func (c *UserController) Logout() {
	ticket := c.GetSession(utils.TicketName)
	//如果为空,则直接返回
	if nil == ticket {
		c.Data["json"] = SuccessVO(nil)
		c.ServeJSON()
		return
	}
	v, _ := utils.RedisDel(ticket.(string))
	logs.Info("del redis:%v", v)
	// utils.RedisDo("del", c.GetSession(utils.TicketName))
	c.DelSession(utils.TicketName)

	c.Data["json"] = SuccessVO(nil)
	c.ServeJSON()
}
