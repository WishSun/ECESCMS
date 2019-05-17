package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseBaseModifyController struct {
	beego.Controller
}

func(this *MyCourseBaseModifyController)Get(){
	this.TplName = "my_course_base_modify.html"

	this.Data["TeacherName"] = this.GetSession("teacher_name")
	tscid := this.Input().Get("tscid")
	this.Data["TSCId"] = tscid
	tsc, err := models.GetTSCBaseById(tscid)
	if err != nil{
		logs.Error(err)
	}
	this.Data["TSC"] = tsc
}

func (this *MyCourseBaseModifyController)Post(){
	tscid := this.Input().Get("TSCId")
	term := this.Input().Get("TSCTerm")
	credit := this.Input().Get("TSCCredit")
	testMethod := this.Input().Get("TSCTestMethod")
	totalPeriod := this.Input().Get("TSCTotalPeriod")
	theoryPeriod := this.Input().Get("TSCTheoryPeriod")
	experimentalPeriod := this.Input().Get("TSCExperimentalPeriod")
	computerPeriod := this.Input().Get("TSCComputerPeriod")
	practicePeriod := this.Input().Get("TSCPracticePeriod")
	weekPeriod := this.Input().Get("TSCWeekPeriod")
	contentRelationImgPath := this.Input().Get("TSCContentRelationImgPath")
	teachTargetOverview := this.Input().Get("TSCTeachTargetOverview")
	classroomTeachTargetOverview := this.Input().Get("TSCClassroomTeachTargetOverview")
	experimentTeachTargetOverview := this.Input().Get("TSCExperimentTeachTargetOverview")
	courseTask := this.Input().Get("TSCCourseTask")
	teachMethod := this.Input().Get("TSCTeachMethod")
	relationOtherCourse := this.Input().Get("TSCRelationOtherCourse")
	category := this.Input().Get("TSCCategory")

	err := models.ModifyTSCBase(tscid, term, credit, testMethod, totalPeriod, theoryPeriod, experimentalPeriod, computerPeriod,
		practicePeriod, weekPeriod, contentRelationImgPath, teachTargetOverview, classroomTeachTargetOverview,
		experimentTeachTargetOverview, courseTask, teachMethod, relationOtherCourse, category)
	if err != nil{
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_base?tscid=%s",tscid), 302)
}
