package controllers

import (
	"ECESCMS/code/common"
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorTrainTargetModifyController struct {
	beego.Controller
}

func (this *MajorTrainTargetModifyController) Get() {
	this.TplName = "major_train_target_modify.html"

	mid := this.Input().Get("mid")
	trts, overview, err := models.GetMajorAllTrainTarget(mid)
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["TrTs"] = trts
		this.Data["Overview"] = overview
		this.Data["Mid"] = mid
	}
}

func (this *MajorTrainTargetModifyController) Post() {
	this.Ctx.Request.ParseForm()

	// reqBody是一个map结构，同name的标签提交后会是一个列表，即map["TrT":[目标1 目标2 目标3], "mid":[2]]
	reqBody := this.Ctx.Request.Form
	mid := this.Input().Get("mid")
	trts := make([]*common.TrTType, 0)

	for i := 0; i < len(reqBody["TrTNumber"]); i++ {
		trts = append(trts, &common.TrTType{reqBody["TrTNumber"][i], reqBody["TrTContent"][i]})
	}

	// 更新专业培养目标
	err := models.ModifyMajorTrainTarget(trts, mid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_train_target?mid=%s", mid), 302)
}
