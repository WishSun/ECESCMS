package main

import (
	"beelog/models"
	_ "beelog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func init() {
	// 注册数据库
	models.RegisterDB()

}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true

	// 创建附件目录
	os.Mkdir("attachment", os.ModePerm)

	// 处理附件方法一 : 作为静态文件
	//beego.SetStaticPath("/attachment", "attachment")

	// 处理附件方法二 : 作为单独一个控制器来处理, 在Router.go中注册下面路由
	// beego.Router("/attachment/:all", &controllers.AttachController{})
	beego.Run()
}
