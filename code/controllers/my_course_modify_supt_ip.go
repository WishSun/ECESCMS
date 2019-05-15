package controllers

import "github.com/astaxie/beego"

type MyCourseModifySuptIPController struct {
	beego.Controller
}

func (this *MyCourseModifySuptIPController) Get() {
	this.TplName = "my_course_modify_supt_ip.html"
}
