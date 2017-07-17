package controllers

import (
	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego/logs"
)

const (
	SuccCode          = "000000"
	SuccMsg           = "success"
	LoginFailCode     = "000100"
	LoginFailMsg      = "登陆失败"
	RegistExistsCode  = "000101"
	RegistExistsMsg   = "用户已注册"
	NoLoginCode       = "000102"
	NoLoginMsg        = "未登录,请先登录"
	ReLoginCode       = "000103"
	ReLoginMsg        = "已登录,请勿重新登陆"
	ParamNotValidCode = "000002"
	ParamNotValidMsg  = "参数校验失败"
	UpdateFailCode    = "000104"
	UpdateFailMsg     = "更新失败"
)

func ValidParams(form interface{}) string {
	valid := validation.Validation{}
	b, err := valid.Valid(form)
	if nil != err {
		logs.Error("err on valid:%v", err)
		return "valid error!"
	}

	var msg = ""
	if !b {
		// validation does not passRegistExistsCode
		// blabla...
		for _, err := range valid.Errors {
			msg = msg + err.Key + ":" + err.Message + "\n\r"
		}

	}
	return msg
}
