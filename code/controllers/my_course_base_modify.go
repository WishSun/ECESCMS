package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"path"
)

type MyCourseBaseModifyController struct {
	beego.Controller
}

func (this *MyCourseBaseModifyController) Get() {
	this.TplName = "my_course_base_modify.html"

	this.Data["TeacherName"] = this.GetSession("teacher_name")
	tscid := this.Input().Get("tscid")
	this.Data["TSCId"] = tscid
	tsc, err := models.GetTSCBaseById(tscid)
	if err != nil {
		logs.Error(err)
	}
	this.Data["TSC"] = tsc
}

func (this *MyCourseBaseModifyController) Post() {
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
	//contentRelationImgPath := this.Input().Get("TSCContentRelationImgPath")
	teachTargetOverview := this.Input().Get("TSCTeachTargetOverview")
	classroomTeachTargetOverview := this.Input().Get("TSCClassroomTeachTargetOverview")
	experimentTeachTargetOverview := this.Input().Get("TSCExperimentTeachTargetOverview")
	courseTask := this.Input().Get("TSCCourseTask")
	teachMethod := this.Input().Get("TSCTeachMethod")
	relationOtherCourse := this.Input().Get("TSCRelationOtherCourse")
	category := this.Input().Get("TSCCategory")

	// 获取图片附件
	_, fh, err := this.GetFile("TSCContentRelationImgPath")
	if err != nil {
		logs.Error(err)
	}
	var contentRelationImgPath string
	if fh != nil {
		// 保存附件
		contentRelationImgPath = tscid + fh.Filename
		logs.Info(contentRelationImgPath)

		// 将附件保存到指定文件夹(这里第二个参数可以是绝对路径，也可以是相对路径)
		// 最后将附件保存在 attachment/附件文件名
		err = this.SaveToFile("TSCContentRelationImgPath", path.Join("attachment", contentRelationImgPath))
		if err != nil {
			logs.Error(err)
		}
	}

	err = models.ModifyTSCBase(tscid, term, credit, testMethod, totalPeriod, theoryPeriod, experimentalPeriod, computerPeriod,
		practicePeriod, weekPeriod, contentRelationImgPath, teachTargetOverview, classroomTeachTargetOverview,
		experimentTeachTargetOverview, courseTask, teachMethod, relationOtherCourse, category)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/teacher/my_course_base?tscid=%s", tscid), 302)
}
