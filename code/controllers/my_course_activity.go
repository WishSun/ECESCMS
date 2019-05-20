package controllers

import (
	"ECESCMS/code/common"
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type MyCourseActivityController struct {
	beego.Controller
}

type MyTeTWeight struct {
	Code       int
	TeTWeights []*common.TeTWeightType
}

type MyTeTs struct {
	Code int
	TeTs []*models.TeachTarget
}

func (this *MyCourseActivityController) Get() {
	this.TplName = "my_course_activity.html"
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	tscid := this.Input().Get("tscid")
	this.Data["TSCId"] = tscid

	tas, err := models.GetTSCTeachActivity(tscid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["TeachActivitys"] = tas
}

// 添加教学活动
func (this *MyCourseActivityController) AddActivity() {
	this.Ctx.Request.ParseForm()
	tetIds := this.Ctx.Request.Form["TeTIds"]
	weights := this.Ctx.Request.Form["Weights"]

	tscid := this.Input().Get("TSCId")
	name := this.Input().Get("ACName")
	weight := this.Input().Get("TAWeight")

	tetws := make([]*common.SendTeTWeight, 0)
	var tetWeightNum int64
	var err error
	for i := 0; i < len(tetIds); i++ {
		if weights[i] == "" {
			tetWeightNum = 0
		} else {
			tetWeightNum, err = strconv.ParseInt(weights[i], 10, 64)
			if err != nil {
				logs.Error(err)
			}
		}
		tetws = append(tetws, &common.SendTeTWeight{TeTId: tetIds[i], Weight: tetWeightNum})
	}

	err = models.AddTSCTeachActivity(tscid, name, weight, tetws)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity?tscid=%s", tscid), 302)
}

// 修改教学活动
func (this *MyCourseActivityController) ModifyActivity() {
	this.Ctx.Request.ParseForm()
	ttwIds := this.Ctx.Request.Form["TTWIds"]
	weights := this.Ctx.Request.Form["Weights"]

	tscid := this.Input().Get("TSCId")
	taid := this.Input().Get("TAId")
	name := this.Input().Get("ACName")
	weight := this.Input().Get("TAWeight")

	tetwms := make([]*common.SendTeTWeightToModify, 0)
	for i := 0; i < len(ttwIds); i++ {
		tetwms = append(tetwms, &common.SendTeTWeightToModify{ttwIds[i], weights[i]})
	}

	err := models.ModifyTSCTeachActivity(taid, name, weight, tetwms)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity?tscid=%s", tscid), 302)
}

// 删除教学活动
func (this *MyCourseActivityController) DeleteActivity() {
	taid := this.Input().Get("taid")
	tscid := this.Input().Get("tscid")
	err := models.DeleteTSCTeachActivity(taid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_activity?tscid=%s", tscid), 302)
}

// 获取教学活动中各个教学目标权重
func (this *MyCourseActivityController) GetTeTWeight() {
	taId := this.Input().Get("TAId")
	m := &MyTeTWeight{}
	tetWs, err := models.GetTATeachTargetWeight(taId)
	if err != nil {
		logs.Error(err)
		m.Code = 1
	} else {
		m.Code = 0
	}
	m.TeTWeights = tetWs

	// 序列化为json后返回
	data, _ := json.Marshal(m)
	fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}

// 获取课程的所有教学目标
func (this *MyCourseActivityController) GetTeTs() {
	tscid := this.Input().Get("TSCId")
	m := MyTeTs{}
	tets, err := models.GetTSCTeachTarget(tscid)
	m.TeTs = tets
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

// 获取携带权值的教学目标
func (this *MyCourseActivityController) GetTeTWs() {
	taid := this.Input().Get("TAId")
	tscid := this.Input().Get("TSCId")

	type Message struct {
		Code     int
		TAName   string
		TAWeight int64
		TeTWs    []*common.ModifyGetTeTWeight
	}

	m := Message{}
	var err error
	mttws := make([]*common.ModifyGetTeTWeight, 0)

	m.TAName, m.TAWeight, mttws, err = models.GetTATeTWs(taid, tscid)
	if err != nil {
		m.Code = 1
		logs.Error(err)
	} else {
		m.Code = 0
	}
	m.TeTWs = mttws

	// 序列化为json后返回
	data, _ := json.Marshal(m)
	fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}
