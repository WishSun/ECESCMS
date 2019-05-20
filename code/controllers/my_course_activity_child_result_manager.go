package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseActivityChildResultManagerController struct {
	beego.Controller
}

func (this *MyCourseActivityChildResultManagerController) Get() {
	this.TplName = "my_course_activity_child_result_manager.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	tacid := this.Input().Get("tacid")
	tacName := this.Input().Get("tacname")
	tacValue := this.Input().Get("tacvalue")
	taid := this.Input().Get("taid")
	this.Data["TAId"] = taid
	this.Data["TACId"] = tacid
	this.Data["TACName"] = tacName
	this.Data["TACValue"] = tacValue
}

func (this *MyCourseActivityChildResultManagerController) Post() {
	this.Ctx.Request.ParseForm()

	sIds := this.Ctx.Request.Form["SIds"]
	sResults := this.Ctx.Request.Form["SResults"]
	tacid := this.Input().Get("TACId")
	tacname := this.Input().Get("TACName")

	err := models.ModifyTACResult(tacid, sIds, sResults)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity_child_result_manager?tacid=%s&tacname=%s", tacid, tacname), 302)
}

// 获取教学活动项成绩
func (this *MyCourseActivityChildResultManagerController) GetTAChildResult() {
	tacid := this.Input().Get("TACId")

	type Message struct {
		Code int
		TACR *models.TAChildResult
	}
	m := Message{}
	var err error
	m.TACR, err = models.GetTAChildResult(tacid)
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
