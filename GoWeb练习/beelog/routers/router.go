package routers

import (
	"beelog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 注册 beego 路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})

	beego.Router("/topic", &controllers.TopicController{})
	// 自动路由，beego只会识别以"Controller"结尾的控制器，这里的"TopicController"明显是可以满足的
	// 然后，它会匹配topic路径下的对应的函数，例如路径"/topic/add", 它就会去找TopicController的Add方法
	// 来处理这个请求，如果没有Add方法的话，就会匹配其他规则，还未找到的话，就会返回404错误
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/reply", &controllers.ReplyController{})
	// 只接受post请求到Add方法中
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	// 只接受get请求到Delete方法中
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	beego.Router("/attachment/:all", &controllers.AttachController{})
}
