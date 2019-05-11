package controllers

import "github.com/astaxie/beego"

type TeacherSelectCourseManagerController struct {
	beego.Controller
}

func (this *TeacherSelectCourseManagerController) Get() {
	this.TplName = "teacher_select_course_manager.html"
}
