package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseBaseController struct {
	beego.Controller
}

func(this*MyCourseBaseController)Get(){
	this.TplName = "my_course_base.html"
	tscid := this.Input().Get("tscid")
	this.Data["TSCId"] = tscid
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	tsc, err := models.GetTSCBaseById(tscid)
	if err != nil{
		logs.Error(err)
	}
	this.Data["TSC"] = tsc
}
