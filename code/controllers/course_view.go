package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CourseViewController struct {
	beego.Controller
}

func (this *CourseViewController) Get() {
	this.TplName = "course_view.html"

	cid := this.Input().Get("cid")
	this.Data["Cid"] = cid

	course, err := models.GetCourse(cid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["Course"] = course
}
