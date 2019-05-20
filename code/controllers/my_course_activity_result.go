package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseActivityResultController struct {
	beego.Controller
}

func (this *MyCourseActivityResultController) Get() {
	this.TplName = "my_course_activity_result.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	taid := this.Input().Get("taid")
	tscid := this.Input().Get("tscid")
	this.Data["TAId"] = taid
	this.Data["TSCId"] = tscid
}

func (this *MyCourseActivityResultController) GetTAResult() {
	type Message struct {
		Code int
		TAR  *models.TAResult
	}

	taid := this.Input().Get("TAId")
	m := Message{}
	var err error
	m.TAR, err = models.GetTAResult(taid)
	m.TAR.TAId = &taid
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
