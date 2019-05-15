package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TeacherSelectCourseManagerController struct {
	beego.Controller
}

func (this *TeacherSelectCourseManagerController) Get() {
	this.TplName = "teacher_select_course_manager.html"

	// 获取所有年级
	grades, err := models.GetAllGrade()
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["Grades"] = grades
	}

	// 获取所有专业
	majors, err := models.GetAllMajor()
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["Majors"] = majors
	}
}

func (this *TeacherSelectCourseManagerController) Post() {
	grade := this.Input().Get("Grade")
	mid := this.Input().Get("Mid")

	this.Redirect(fmt.Sprintf("/admin/teacher_select_course_view?grade=%s&mid=%s", grade, mid), 302)
}
