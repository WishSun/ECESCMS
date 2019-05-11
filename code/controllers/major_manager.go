package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorManagerController struct {
	beego.Controller
}

func (this *MajorManagerController) Get() {
	this.TplName = "major_manager.html"

	majors, err := models.GetAllMajor()
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["Majors"] = majors
	}
}

func (this *MajorManagerController) Delete() {
	mid := this.Input().Get("mid")

	err := models.DeleteMajor(mid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect("/admin/major_manager", 302)
}
