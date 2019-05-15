package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorCourseController struct {
	beego.Controller
}

func (this *MajorCourseController) Get() {
	this.TplName = "major_course.html"

	mid := this.Input().Get("mid")
	this.Data["Mid"] = mid

	// 获取专业所有课程
	courses, err := models.GetMajorAllCourse(mid)
	if err != nil {
		logs.Error(err)
	}

	this.Data["Courses"] = courses
}
