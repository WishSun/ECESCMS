package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseTeachTargetController struct {
	beego.Controller
}

type GRIPType struct {
	Code int
	GRs  []*models.GraduationRequirement
	IPs  []*models.IndicatorPoint
}

func (this *MyCourseTeachTargetController) Get() {
	this.TplName = "my_course_teach_target.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	this.Data["TSCId"] = this.Input().Get("tscid")

	TeTs, err := models.GetTSCTeachTarget(this.Input().Get("tscid"))
	if err != nil {
		logs.Error(err)
	}
	this.Data["TeTs"] = TeTs
}

// 获取毕业要求和指标点
func (this *MyCourseTeachTargetController) GetGRAndIP() {
	tetId := this.Input().Get("TeTId")

	grs, ips, err := models.GetTSCTeachTargetGRAndIP(tetId)
	var grips *GRIPType
	if err != nil {
		logs.Error(err)
		grips = &GRIPType{1, grs, ips}
	} else {
		grips = &GRIPType{0, grs, ips}
	}

	data, _ := json.Marshal(grips)
	fmt.Println(this.Ctx.ResponseWriter, string(data))
}

// 添加教学目标
func (this *MyCourseTeachTargetController) AddTeachTarget() {
	number := this.Input().Get("TeTNumber")
	description := this.Input().Get("TeTDescription")
	tscid := this.Input().Get("TSCId")

	err := models.AddTeachTarget(tscid, number, description)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_teach_target?tscid=%s", tscid), 302)
}

// 修改教学目标
func (this *MyCourseTeachTargetController) ModifyTeachTarget() {
	number := this.Input().Get("TeTNumber")
	description := this.Input().Get("TeTDescription")
	tetid := this.Input().Get("TeTId")
	tscid := this.Input().Get("TSCId")

	err := models.ModifyTeachTarget(tetid, number, description)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_teach_target?tscid=%s", tscid), 302)
}

// 删除教学目标
func (this *MyCourseTeachTargetController) DeleteTeachTarget() {
	err := models.DeleteTeachTarget(this.Input().Get("tetid"))
	if err != nil {
		logs.Error(err)
	}
	tscid := this.Input().Get("tscid")

	this.Redirect(fmt.Sprintf("/teacher/my_course_teach_target?tscid=%s", tscid), 302)
}
