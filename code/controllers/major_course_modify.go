package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorCourseModifyController struct {
	beego.Controller
}

func (this *MajorCourseModifyController) Get() {
	this.TplName = "major_course_modify.html"

	mid := this.Input().Get("mid")
	this.Data["Mid"] = mid

	// 获取所有课程
	courses, err := models.GetAllCourse()
	if err != nil {
		logs.Error(err)
	}

	this.Data["Courses"] = courses
}

func (this *MajorCourseModifyController) Post() {
	err := this.Ctx.Request.ParseForm()
	if err != nil {
		logs.Error(err)
	}
	mid := this.Input().Get("Mid")

	cids := this.Ctx.Request.Form["MCourses"]
	err = models.ModifyMajorCourses(mid, cids)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_course?mid=%s", mid), 302)
}
