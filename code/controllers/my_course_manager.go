package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseManagerController struct {
	beego.Controller
}

func (this *MyCourseManagerController) Get() {
	this.TplName = "my_course_manager.html"

	this.Data["TeacherName"] = this.GetSession("teacher_name")
	var tid int64
	tid = this.GetSession("teacher_id").(int64)
	courseList, err := models.GetTeacherCourseList(tid)
	if err != nil{
		logs.Error(err)
	}
	this.Data["CourseList"] = courseList
}
