package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CourseAddController struct {
	beego.Controller
}

func (this *CourseAddController) Get() {
	this.TplName = "course_add.html"
}

func (this *CourseAddController) Post() {
	number := this.Input().Get("CourseNumber")
	nameCN := this.Input().Get("CourseNameCN")
	nameEN := this.Input().Get("CourseNameEN")
	knowledgeField := this.Input().Get("CourseKnowledgeField")
	usedMajors := this.Input().Get("CourseUsedMajors")
	preparatoryCourse := this.Input().Get("CoursePreparatoryCourse")
	recommendTextBook := this.Input().Get("CourseRecommendTextBook")
	referenceBook := this.Input().Get("CourseReferenceBook")
	relativeWebsite := this.Input().Get("CourseRelativeWebsite")

	err := models.AddCourse(number, nameCN, nameEN, knowledgeField, usedMajors, preparatoryCourse, recommendTextBook, referenceBook, relativeWebsite)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/course_manager", 302)
}
