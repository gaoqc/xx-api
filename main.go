package main

import (
	_ "xx-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"xx-api/models"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initDb()
	initLog()
	beego.Run()

}
func initDb() {
	beego.Debug("begin initDB")
	orm.RegisterDataBase("default", "mysql", "root:gaoqc@123.com@/xx?charset=utf8")
	orm.RegisterModelWithPrefix("t_",
		new(models.User), new(models.UserAddress), new(models.Article), new(models.CommentLike),
		new(models.Comment), new(models.CustOrder), new(models.Province), new(models.City), new(models.Area),
		new(models.Street), new(models.Vendor), new(models.HomeAppKind))
	orm.Debug = true
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		beego.Error("err on RunSyncdb!")

	}
	beego.Debug("end initDB")
}

func initLog() {
	logs.SetLevel(beego.LevelDebug)
	logs.SetLogger("console")
	// beego.SetLogger("file", `{"filename":"/Users/gaoqc/Documents/codes/go/src/xx-api/logs/xx-api.log"}`)

}
