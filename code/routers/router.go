package routers

import (
	"ECESCMS/code/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
