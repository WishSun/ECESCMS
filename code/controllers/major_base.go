package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorBaseController struct {
	beego.Controller
}

func (this *MajorBaseController) Get() {
	this.TplName = "major_base.html"

	mid := this.Input().Get("mid")
	this.Data["Mid"] = mid

	major, err := models.GetMajorBase(mid)
	if err != nil {
		logs.Error(err)
	}

	this.Data["Major"] = major
}
