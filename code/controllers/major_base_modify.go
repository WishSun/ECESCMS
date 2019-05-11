package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorBaseModifyController struct {
	beego.Controller
}

func (this *MajorBaseModifyController) Get() {
	this.TplName = "major_base_modify.html"

	mid := this.Input().Get("mid")
	major, err := models.GetMajorBase(mid)
	if err != nil {
		logs.Error(err)
	}

	this.Data["Major"] = major
}

func (this *MajorBaseModifyController) Post() {
	// 解析表单
	mid := this.Input().Get("mid")
	MajorNumber := this.Input().Get("MajorNumber")
	MajorName := this.Input().Get("MajorName")
	TrTOverview := this.Input().Get("TrTOverview")
	MainSubject := this.Input().Get("MainSubject")
	StudyYears := this.Input().Get("StudyYears")
	Degree := this.Input().Get("Degree")
	CoreCourse := this.Input().Get("CoreCourse")
	TotalCredits := this.Input().Get("TotalCredits")
	TheoryCredits := this.Input().Get("TheoryCredits")
	PracticeCredits := this.Input().Get("PracticeCredits")

	// 调用修改方法
	err := models.ModifyMajorBase(mid, MajorNumber, MajorName, TrTOverview, MainSubject, Degree, CoreCourse,
		StudyYears, TotalCredits, TheoryCredits, PracticeCredits)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_base?mid=%s", mid), 302)
}
