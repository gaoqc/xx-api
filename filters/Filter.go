package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	// "strconv"
	"strings"
	"xx-api/controllers"
	"xx-api/utils"
)

var PrintURL = func(ctx *context.Context) {
	// uri := ctx.Input.URI()
	// beego.Info("uri:" + uri)
	// for k, v := range ctx.Input.Params() {
	// 	beego.Info(k + "=" + v)
	// }
}

/**
* 跨域设置
**/

var CorsAllow = func(ctx *context.Context) {
	// allowurl := "http://127.0.0.1:8080"
	allowurl := "http://localhost:8080"
	logs.Info("allow url:%s", allowurl)

	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST")
	ctx.Output.Header("Access-Control-Allow-Origin", allowurl)
	ctx.Output.Header("Access-Control-Expose-Headers", "Content-Length")
}

/**
* 登陆控制区
**/
func loginFilterUri() []string {
	return []string{"/v1/user/update", "/v1/userAddress", "/v1/user/chgPwd",
		"/v1/user/SendChgPwdValidCode", "/v1/article/add", "/v1/article/del",
		"/v1/article/myList", "/v1/article/comment/add", "/v1/article/comment/del",
		"/v1/article/comment/like/", "/v1/cust/addOrder", "/v1/cust/qryOrderList"}
}

var LoginFilter = func(ctx *context.Context) {
	uri := ctx.Input.URI()
	ticket := ctx.Input.Session(utils.TicketName)
	logs.Debug("uri is:" + uri + " ,ticket is :" + utils.ToJson(ticket))
	for _, f := range loginFilterUri() {
		if strings.HasPrefix(uri, f) {
			beego.Debug("uri:" + uri + ",patter:" + f)
			if ticket == nil {
				beego.Warn("uri:" + uri + " not login")
				ctx.Output.JSON(controllers.GetRetVO(controllers.NoLoginCode, controllers.NoLoginMsg, nil), false, false)
			}
			break
		}
	}

}
