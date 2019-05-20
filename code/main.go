package main

import (
	"ECESCMS/code/models"
	_ "ECESCMS/code/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}
func main() {
	// 设置开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	// 设置日志显示文件和行号
	logs.SetLogFuncCall(true)
	orm.Debug = true
	beego.Run()
}
