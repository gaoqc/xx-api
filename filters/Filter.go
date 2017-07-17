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

func loginFilterUri() []string {
	return []string{"/v1/user/update", "/v1/userAddress"}
}

var LoginFilter = func(ctx *context.Context) {
	session := beego.BConfig.WebConfig.Session
	logs.Debug("SessionGCMaxLifetime:%d,cookie'time:%d", session.SessionGCMaxLifetime, session.SessionCookieLifeTime)
	uri := ctx.Input.URI()
	ticket := ctx.Input.Session(utils.TicketName)
	logs.Debug("uri is" + uri + " ,ticket is :" + utils.ToJson(ticket))
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
