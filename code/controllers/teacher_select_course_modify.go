package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TeacherSelectCourseModifyController struct {
	beego.Controller
}

func (this *TeacherSelectCourseModifyController) Get() {
	this.TplName = "teacher_select_course_modify.html"

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

	teachers, err := models.GetAllTeacher()
	if err != nil {
		logs.Error(err)
	}
	this.Data["Teachers"] = teachers
}

func (this *TeacherSelectCourseModifyController) Post() {
	_ = this.Ctx.Request.ParseForm()
	mid := this.Input().Get("Mid")
	grade := this.Input().Get("Grade")
	CourseIds := this.Ctx.Request.Form["CourseId"]
	TeacherInfos := this.Ctx.Request.Form["TeacherInfo"]

	err := models.ModifyMajorTSC(mid, grade, CourseIds, TeacherInfos)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/teacher_select_course_view?mid=%s&grade=%s", mid, grade), 302)
}
