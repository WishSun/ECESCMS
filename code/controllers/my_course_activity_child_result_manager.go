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
	tscid := this.Input().Get("tscid")
	this.Data["TAId"] = taid
	this.Data["TACId"] = tacid
	this.Data["TACName"] = tacName
	this.Data["TACValue"] = tacValue
	this.Data["TSCId"] = tscid

	tets, err := models.GetTASuptTeachTargets(taid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["TeTs"] = tets
}

func (this *MyCourseActivityChildResultManagerController) Post() {
	this.Ctx.Request.ParseForm()

	sIds := this.Ctx.Request.Form["SIds"]
	sResults := this.Ctx.Request.Form["SResults"]
	tacid := this.Input().Get("TACId")
	tacName := this.Input().Get("TACName")
	taid := this.Input().Get("TAId")
	tacValue := this.Input().Get("TACValue")
	tscid := this.Input().Get("TSCId")
	this.Data["TAId"] = taid
	this.Data["TACId"] = tacid
	this.Data["TACName"] = tacName
	this.Data["TACValue"] = tacValue
	this.Data["TSCId"] = tscid

	err := models.ModifyTACResult(tacid, sIds, sResults)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity_child_result_manager"+
		"?tacid=%s&tacname=%s&taid=%s&tscid=%s&tacvalue=%s", tacid, tacName, taid, tscid, tacValue), 302)
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

// 获取活动项信息
func (this *MyCourseActivityChildResultManagerController) GetTAChildInfo() {
	tacid := this.Input().Get("TACId")

	type Message struct {
		Code     int
		TAChild  *models.TeachActivityChild
		SuptTeTs []*models.TACSuptTeT
	}
	m := Message{}
	var err1, err2 error
	m.TAChild, err1 = models.GetTeachActivityChild(tacid)
	m.SuptTeTs, err2 = models.GetTACSuptTeachTargets(tacid)
	if err1 != nil && err2 != nil {
		logs.Error(err1)
		logs.Error(err2)
		m.Code = 1
	} else {
		m.Code = 0
	}

	// 序列化为json后返回
	data, _ := json.Marshal(m)
	fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}

// 删除教学活动项
func (this *MyCourseActivityChildResultManagerController) DeleteTAChild() {
	taid := this.Input().Get("taid")
	tscid := this.Input().Get("tscid")
	tacid := this.Input().Get("tacid")
	err := models.DeleteTSCTeachActivityChild(tacid)
	if err != nil {
		logs.Error(err)
	}
	this.Redirect(fmt.Sprintf("/teacher/my_course_activity_child?taid=%s&tscid=%s", taid, tscid), 302)
}

// 修改教学活动项
func (this *MyCourseActivityChildResultManagerController) ModifyTAChild() {
	tacid := this.Input().Get("TACId")
	tacname := this.Input().Get("TAChildName")
	tacvalue := this.Input().Get("TAChildValue")
	tscid := this.Input().Get("TSCId")
	taid := this.Input().Get("TAId")

	this.Ctx.Request.ParseForm()
	tetIds := this.Ctx.Request.Form["TeTIds"]

	err := models.ModifyTSCTeachActivityChild(tacid, tacname, tacvalue, tetIds)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity_child_result_manager?tacid=%s&taid=%s&tscid=%s&tacname=%s&tacvalue=%s", tacid, taid, tscid, tacname, tacvalue), 302)
}
