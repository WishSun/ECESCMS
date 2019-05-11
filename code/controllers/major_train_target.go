package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorTrainTargetController struct {
	beego.Controller
}

func (this *MajorTrainTargetController) Get() {
	this.TplName = "major_train_target.html"

	mid := this.Input().Get("mid")
	this.Data["Mid"] = mid

	trts, overview, err := models.GetMajorAllTrainTarget(mid)
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["TrTs"] = trts
		this.Data["Overview"] = overview
	}
}
