package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseResultAnalysisReportController struct {
	beego.Controller
}

func (this *MyCourseResultAnalysisReportController) Get() {
	this.TplName = "my_course_result_analysis_report.html"

	tscid := this.Input().Get("tscid")
	tets, err := models.GetTSCTeachTarget(tscid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["TeTs"] = tets
	this.Data["TSCId"] = tscid

	course, tsc, err := models.GetCourseAndTSC(tscid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["TSC"] = tsc
	this.Data["Course"] = course

	analysisReport, err := models.GetTSCResultAnalysisReport(tscid)
	this.Data["AnalysisReport"] = analysisReport
}

func (this *MyCourseResultAnalysisReportController) Post() {
	pt := this.Input().Get("PTText")
	rep := this.Input().Get("REPTText")
	aesn := this.Input().Get("AESNText")
	el := this.Input().Get("ELText")
	it := this.Input().Get("ITText")

	tscid := this.Input().Get("TSCId")

	tetim := this.Input().Get("TeTImprovementMeasures")
	tetfa := this.Input().Get("TeTFinishAnalysis")
	gripim := this.Input().Get("GRIPImprovementMeasures")
	id := this.Input().Get("Id")

	err := models.ModifyTSCResultAnalysisReport(id, pt, rep, aesn, el, it, tetim, gripim, tetfa)
	if err != nil {
		logs.Error(err)
	}
	this.Redirect(fmt.Sprintf("/teacher/my_course_result_analysis_report?tscid=%s", tscid), 302)
}

func (this *MyCourseResultAnalysisReportController) GetActivityInfo() {
	type Message struct {
		Code         int
		AnalysisData *models.ResultAnalysisData
	}

	tscid := this.Input().Get("TSCId")
	m := &Message{}
	var err error
	m.AnalysisData, err = models.GetResultAnalysisData(tscid)
	if err != nil {
		logs.Error(err)
	}

	data, _ := json.Marshal(m)
	_, _ = fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}
