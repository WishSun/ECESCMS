package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CourseModifyController struct {
	beego.Controller
}

func (this *CourseModifyController) Get() {
	this.TplName = "course_modify.html"

	cid := this.Input().Get("cid")
	this.Data["Cid"] = cid

	course, err := models.GetCourse(cid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["Course"] = course
}

func (this *CourseModifyController) Post() {
	cid := this.Input().Get("Cid")
	number := this.Input().Get("CourseNumber")
	nameCN := this.Input().Get("CourseNameCN")
	nameEN := this.Input().Get("CourseNameEN")
	knowledgeField := this.Input().Get("CourseKnowledgeField")
	usedMajors := this.Input().Get("CourseUsedMajors")
	preparatoryCourse := this.Input().Get("CoursePreparatoryCourse")
	recommendTextBook := this.Input().Get("CourseRecommendTextBook")
	referenceBook := this.Input().Get("CourseReferenceBook")
	relativeWebsite := this.Input().Get("CourseRelativeWebsite")

	err := models.ModifyCourse(cid, number, nameCN, nameEN, knowledgeField, usedMajors, preparatoryCourse, recommendTextBook, referenceBook, relativeWebsite)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/course_view?cid=%s", cid), 302)
}
