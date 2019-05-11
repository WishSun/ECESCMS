package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

	// 在登录成功后，在session中设置登录对象
	this.SetSession("obj", this.Input().Get("obj"))

	IsAdmin := this.GetSession("obj") == "admin"
	this.Data["IsAdmin"] = IsAdmin
}

func (this *LoginController) Post() {
	if this.GetSession("obj") == "admin" {
		this.Redirect("/admin_home", 301)
	} else {
		this.Redirect("teacher_home", 301)
	}
}
