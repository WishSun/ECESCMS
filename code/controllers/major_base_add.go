package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type MajorBaseAddController struct {
	beego.Controller
}

func (this *MajorBaseAddController) Get() {
	this.TplName = "major_base_add.html"

	courses, err := models.GetAllCourse()
	if err != nil {
		logs.Error(err)
	}

	this.Data["Courses"] = courses
}

func (this *MajorBaseAddController) Post() {
	// 解析表单
	MajorNumber := this.Input().Get("MajorNumber")
	MajorName := this.Input().Get("MajorName")
	TrTOverview := this.Input().Get("TrTOverview")
	MainSubject := this.Input().Get("MainSubject")
	StudyYears, _ := strconv.ParseInt(this.Input().Get("StudyYears"), 10, 64)
	Degree := this.Input().Get("Degree")
	CoreCourse := this.Input().Get("CoreCourse")
	TotalCredits, _ := strconv.ParseFloat(this.Input().Get("TotalCredits"), 64)
	TheoryCredits, _ := strconv.ParseFloat(this.Input().Get("TheoryCredits"), 64)
	PracticeCredits, _ := strconv.ParseFloat(this.Input().Get("PracticeCredits"), 64)

	// 调用添加专业方法
	err := models.AddMajor(MajorNumber, MajorName, TrTOverview, MainSubject, Degree, CoreCourse,
		StudyYears, TotalCredits, TheoryCredits, PracticeCredits)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/major_manager", 302)
}
