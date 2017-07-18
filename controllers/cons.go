package controllers

import (
	"github.com/astaxie/beego/validation"

	"github.com/astaxie/beego/logs"
)

const (
	/**
	 000000-100000 定义公共异常码
	**/

	SuccCode          = "000000"
	SuccMsg           = "success"
	ParamNotValidCode = "000002"
	ParamNotValidMsg  = "参数校验失败"
	ParamNotEmptyCode = "000003"
	ParamNotEmptyMsg  = "参数%s不能为空"
	UpdateFailCode    = "000104"
	UpdateFailMsg     = "更新失败,记录不存在"

	/**
	**100001-101000 定义用户相关异常
	**/
	LoginFailCode     = "100001"
	LoginFailMsg      = "登陆失败,用户名或密码错误"
	RegistExistsCode  = "100002"
	RegistExistsMsg   = "用户已注册"
	NoLoginCode       = "100003"
	NoLoginMsg        = "未登录,请先登录"
	ReLoginCode       = "100004"
	ReLoginMsg        = "已登录,请勿重新登陆"
	ChgPwdInValidCode = "100005"
	ChgPwdInValidMsg  = "验证码校验失败"
	/**
	**1010001-1020001 定义论坛相关异常
	**/
	CommentReLikeCode = "1010001"
	CommentReLikeMsg  = "已点赞,请勿重复点"
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
