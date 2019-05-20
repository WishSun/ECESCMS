package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseActivityChildStudentResultManagerController struct {
	beego.Controller
}

func (this *MyCourseActivityChildStudentResultManagerController) Get() {
	this.TplName = "my_course_activity_child_student_result_manager.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	sid := this.Input().Get("sid")
	taid := this.Input().Get("taid")
	tscid := this.Input().Get("tscid")

	this.Data["TSCId"] = tscid
	this.Data["SId"] = sid
	this.Data["TAId"] = taid
}

func (this *MyCourseActivityChildStudentResultManagerController) Post() {
	this.Ctx.Request.ParseForm()

	sid := this.Input().Get("SId")
	taid := this.Input().Get("TAId")
	tacIds := this.Ctx.Request.Form["TACIds"]
	stacResults := this.Ctx.Request.Form["STACResults"]

	err := models.ModifyStudentTACResults(sid, tacIds, stacResults)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity_child_student_result_manager?sid=%s&taid=%s", sid, taid), 302)
}

func (this *MyCourseActivityChildStudentResultManagerController) GetStudentAllTAChildResult() {
	taid := this.Input().Get("TAId")
	sid := this.Input().Get("SId")

	type Message struct {
		Code   int
		STACRs *models.STAChildResults
	}
	m := Message{}
	var err error
	m.STACRs, err = models.GetSTAChildResults(sid, taid)
	if err != nil {
		logs.Error(err)
		m.Code = 1
	} else {
		m.Code = 0
	}

	// 序列化为json后返回
	data, _ := json.Marshal(m)
	fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}
