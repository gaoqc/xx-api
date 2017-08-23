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
	// "github.com/astaxie/beego/plugins/cors"
	"xx-api/controllers"
	"xx-api/filters"
)

func init() {
	initFilter()
	// nsIndex := beego.NewNamespace("/1", beego.NSNamespace("/",
	// 	beego.NSInclude(
	// 		&controllers.IndexController{},
	// 	),
	// ))
	beego.Include(&controllers.IndexController{})
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
		beego.NSNamespace("/article/comment/like",
			beego.NSInclude(
				&controllers.CommentLikeController{},
			),
		),
		beego.NSNamespace("/cust/",
			beego.NSInclude(
				&controllers.CustOrderController{},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.HomeAppsController{},
			),
		),
		beego.NSNamespace("/vendor",
			beego.NSInclude(
				&controllers.VendorController{},
			),
		),
	)
	// beego.AddNamespace(nsIndex)
	beego.AddNamespace(ns)

}
func initFilter() {

	beego.InsertFilter("/*", beego.BeforeRouter, filters.CorsAllow)
	// 	cors.Allow(&cors.Options{
	// 	AllowAllOrigins:  true,
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{"PUT", "GET", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// })
	// )
}
