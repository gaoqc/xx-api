package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
	"xx-api/models"
)

const (
	Ticket = "ticket"
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
		c.Data["json"] = vo.GetRetVO(vo.RegistExistsCode, vo.RegistExistsMsg, nil)
		c.ServeJSON()
		return
	}

	//新增记录
	id := models.AddUser(&models.User{LoginAcc: userForm.LoginAcc, LoginPwd: userForm.LoginPwd, Phone: userForm.Phone, Email: userForm.Email, TrueName: userForm.TrueName, IdNo: userForm.IdNo, IdType: userForm.IdType, UserType: userForm.UserType})
	c.Data["json"] = vo.SuccessVO("id =" + strconv.FormatInt(id, 10))
	c.ServeJSON()

}
func (c *UserController) Login() {
	loginAcc := c.GetString("loginAcc")
	loginPwd := c.GetString("loginPwd")
	exist, user := models.Login(loginAcc, loginPwd)
	if exist {
		c.SetSession(Ticket, &user)
		c.Data["json"] = vo.SuccessVO(user)
		// beego.Info("ticket is :" + utils.ToJson(c.GetSession(Ticket)))
	} else {
		c.Data["json"] = vo.GetRetVO(vo.LoginFailCode, vo.LoginFailMsg, nil)
	}
	c.ServeJSON()

}

//退出登陆
func (c *UserController) Logout() {
	c.DelSession(Ticket)
	c.Data["json"] = vo.SuccessVO(nil)
	c.ServeJSON()
}
