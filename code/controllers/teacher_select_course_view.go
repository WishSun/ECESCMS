package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TeacherSelectCourseViewController struct {
	beego.Controller
}

func (this *TeacherSelectCourseViewController) Get() {
	this.TplName = "teacher_select_course_view.html"

	mid := this.Input().Get("mid")
	grade := this.Input().Get("grade")

	this.Data["Mid"] = mid
	major, err := models.GetMajorBase(mid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["MajorName"] = major.Name
	this.Data["Grade"] = grade
	tscs, err := models.GetMajorAllSelectCourse(mid)
	if err != nil {
		logs.Error(err)
	}

	this.Data["TSCs"] = tscs
}
