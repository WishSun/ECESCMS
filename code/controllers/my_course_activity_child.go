package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseActivityChildController struct {
	beego.Controller
}

func (this *MyCourseActivityChildController) Get() {
	this.TplName = "my_course_activity_child.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")

	taid := this.Input().Get("taid")
	this.Data["TAId"] = taid

	tscid := this.Input().Get("tscid")
	this.Data["TSCId"] = tscid
	tets, err := models.GetTASuptTeachTargets(taid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["TeTs"] = tets
}

func (this *MyCourseActivityChildController) Post() {
	this.Ctx.Request.ParseForm()

	tscid := this.Input().Get("TSCId")
	taid := this.Input().Get("TAId")
	name := this.Input().Get("AChildName")
	value := this.Input().Get("AChildValue")

	// 选中的教学目标的Id列表
	tetIds := this.Ctx.Request.Form["TeTIds"]

	err := models.AddTSCTeachActivityChild(taid, name, value, tetIds)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity_child?taid=%s&tscid=%s", taid, tscid), 302)
}

// 获取所有活动子项成绩表
func (this *MyCourseActivityChildController) GetTAChildResults() {
	type Message struct {
		Code  int
		TACRs *models.TAChildResults
	}

	taid := this.Input().Get("TAId")
	m := Message{}
	var err error
	m.TACRs, err = models.GetTAChildResults(taid)
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
