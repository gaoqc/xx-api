// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"xx-api/controllers"
	"xx-api/filters"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(filters.LoginFilter),
		beego.NSNamespace("/userAddress",
			beego.NSInclude(
				&controllers.UserAddressController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
		beego.NSNamespace("/article/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
func initFilter() {

	beego.InsertFilter("/*", beego.BeforeRouter, filters.PrintURL)
}
