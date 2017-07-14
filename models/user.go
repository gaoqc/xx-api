package models

import (
	// "github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	TimeModel
	LoginAcc string
	LoginPwd string
	Phone    string `orm:"null"`
	Email    string `orm:"null`
	TrueName string `orm:"null`
	IdNo     string `orm:"null`
	IdType   int    `orm:"null`
	//用户类型,1普通顾客,2表示维修人员
	UserType int `orm:"default(1)"`
	//0 正常,1锁定
	Status int `orm:"default(0)"`
}

//新增记录
func AddUser(user *User) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if err != nil {
		beego.Error("Insert err:")

	}
	return id

}

//更新手机号码或者邮箱或者证件ID.
func Update(phone, email, idNo, trueName, loginAcc string, idType, userType int) int64 {
	user := User{Phone: phone, Email: email, IdNo: idNo, IdType: idType, TrueName: trueName}
	up := make(map[string]interface{})

	if len(phone) > 0 {
		up["phone"] = phone

	}
	if len(email) > 0 {
		up["email"] = email

	}
	if len(trueName) > 0 {
		up["trueName"] = trueName
	}
	if len(idNo) > 0 {
		up["id_no"] = idNo
		up["id_type"] = idType

	}
	if userType > 0 {
		up["user_type"] = userType
	}

	num, err := orm.NewOrm().QueryTable(user).Filter("loginAcc", loginAcc).Update(up)
	if err != nil {
		beego.Error("err at update user")
	}
	return num
}

//登录
func Login(loginAcc, loginPwd string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("login_acc", loginAcc).Filter("login_pwd", loginPwd).One(&user)
	return err != orm.ErrNoRows, user
}

//判断用户是否已注册
func ExistUser(loginAcc, idNo, phone, email string) bool {
	cond := orm.NewCondition()

	qs := orm.NewOrm().QueryTable(User{})

	if len(loginAcc) > 0 {
		cond = cond.Or("login_acc", loginAcc)
	}
	if len(idNo) > 0 {
		cond = cond.Or("id_no", idNo)

	}
	if len(phone) > 0 {
		cond = cond.Or("phone", phone)

	}
	if len(email) > 0 {
		cond = cond.Or("email", email)

	}

	return qs.SetCond(cond).Exist()
}
