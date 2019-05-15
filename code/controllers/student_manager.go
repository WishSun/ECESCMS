package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type StudentManagerController struct {
	beego.Controller
}

func (this *StudentManagerController) Get() {
	this.TplName = "student_manager.html"
	logs.Info("进入Get！")
	majors, err := models.GetAllMajor()
	if err != nil {
		logs.Error(err)
	}
	this.Data["Majors"] = majors

	action := this.Input().Get("action")
	if action == "by_name" {
		name := this.Input().Get("sname")
		students, err := models.GetStudentByName(name)
		if err != nil {
			logs.Error(err)
		}
		logs.Info("学生信息:", students)
		this.Data["Students"] = students
	} else if action == "by_other" {
		majorName := this.Input().Get("mname")
		grade := this.Input().Get("sgrade")
		class := this.Input().Get("sclass")

		// 根据专业名、年级、班级来筛选学生
		students, err := models.GetAllStudent(majorName, grade, class)
		if err != nil {
			logs.Error(err)
		}
		this.Data["Students"] = students
	}
}

func (this *StudentManagerController) Post() {
	action := this.Input().Get("action")
	//this.Redirect(fmt.Sprintf("/admin/student_manager?action=%s", action), 302)
	//

	if action == "by_name" {
		name := this.Input().Get("FilterSName")
		this.Redirect(fmt.Sprintf("/admin/student_manager?action=by_name&sname=%s", name), 302)
	} else if action == "by_other" {
		majorName := this.Input().Get("MNameFilter")
		grade := this.Input().Get("GradeFilter")
		class := this.Input().Get("ClassFilter")
		this.Redirect(fmt.Sprintf("/admin/student_manager?action=by_other&mname=%s&sgrade=%s&sclass=%s",
			majorName, grade, class), 302)
	}

}

func (this *StudentManagerController) Add() {
	number := this.Input().Get("SNumber")
	name := this.Input().Get("SName")
	majorName := this.Input().Get("SMajorName")
	class := this.Input().Get("SClass")
	grade := this.Input().Get("SGrade")

	err := models.AddStudent(number, name, class, grade, majorName)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/student_manager", 302)
}

func (this *StudentManagerController) Modify() {
	sid := this.Input().Get("SId")
	number := this.Input().Get("SNumber")
	name := this.Input().Get("SName")
	class := this.Input().Get("SClass")
	grade := this.Input().Get("SGrade")
	majorName := this.Input().Get("SMajorName")

	err := models.ModifyStudent(sid, number, name, class, grade, majorName)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/student_manager?action=by_name&sname=%s", name), 302)
}

func (this *StudentManagerController) Delete() {
	sid := this.Input().Get("sid")

	err := models.DeleteStudent(sid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/student_manager", 302)
}
