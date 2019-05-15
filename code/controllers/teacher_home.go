package controllers

import "github.com/astaxie/beego"

type TeacherHomeController struct {
	beego.Controller
}

func (this *TeacherHomeController) Get() {
	this.TplName = "teacher_home.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")

}
