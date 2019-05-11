package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CourseManagerController struct {
	beego.Controller
}

func (this *CourseManagerController) Get() {
	this.TplName = "course_manager.html"

	// 获取所有课程
	courses, err := models.GetAllCourse()
	if err != nil {
		logs.Error(err)
	}

	this.Data["Courses"] = courses
}

// 删除课程
func (this *CourseManagerController) Delete() {
	cid := this.Input().Get("cid")

	err := models.DeleteCourse(cid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/course_manager", 302)
}
