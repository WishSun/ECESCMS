package controllers

import (
	"ECESCMS/code/models"
	"github.com/astaxie/beego"
)

type MajorGRIPController struct {
	beego.Controller
}

func (this *MajorGRIPController) Get() {
	this.TplName = "major_GR_IP.html"

	grid := this.Input().Get("grid")
	this.Data["GRId"] = grid

	// 获取该毕业要求的所有指标点
	ips, _ := models.GetMajorGRAllIndicatorPoint(grid)
	this.Data["Ips"] = ips
}
