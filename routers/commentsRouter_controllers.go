package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["xx-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xx-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xx-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xx-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xx-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Myist",
			Router: `/myList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:ArticleController"] = append(beego.GlobalControllerRouter["xx-api/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"] = append(beego.GlobalControllerRouter["xx-api/controllers:CommentLikeController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "AddAddress",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "ListAddress",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserAddressController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserAddressController"],
		beego.ControllerComments{
			Method: "DelAddress",
			Router: `/del`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "SendChgPwdValidCode",
			Router: `/SendChgPwdValidCode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "ChgPassword",
			Router: `/chgPwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xx-api/controllers:UserController"] = append(beego.GlobalControllerRouter["xx-api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
