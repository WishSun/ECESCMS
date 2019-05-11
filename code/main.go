package main

import (
	"ECESCMS/code/models"
	_ "ECESCMS/code/routers"
	"github.com/astaxie/beego"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}
func main() {
	// 设置开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
