package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

// 获取登录页面，即GET请求
func (this *LoginController) Get() {
	// 判断用户是否携带退出参数
	isExit := this.Input().Get("quit") == "true"

	if isExit {
		// 设置cookie为空，即删除之前的cookie, cookie过期时间设置为-1就是立即删除的效果
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")

		// 重定向到首页
		this.Redirect("/", 301)

		return
	}
	this.TplName = "login.html"
}

// 登录信息的获取，即Post请求
func (this *LoginController) Post() {
	// 获取登录信息
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	// 如果账号密码验证成功
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0

		// 设置cookie过期时间
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		// 设置cookie
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	// 重定向
	this.Redirect("/", 301)
	return
}

// 解析是否携带cookie，并检验cookie信息
func checkAccount(ctx *context.Context) bool {
	logs.Info("验证cookie信息")

	// 获取请求所携带的cookie信息(账户和密码)
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	// 验证cookie信息是否正确
	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}
