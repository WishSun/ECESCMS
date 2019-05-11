package controllers

import "github.com/astaxie/beego"

type AdminHomeController struct {
	beego.Controller
}

func (this *AdminHomeController) Get() {
	this.TplName = "admin_home.html"
}
