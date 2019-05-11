package controllers

import (
	"ECESCMS/code/common"
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorGRIPModifyController struct {
	beego.Controller
}

func (this *MajorGRIPModifyController) Get() {
	this.TplName = "major_GR_IP_modify.html"
	grid := this.Input().Get("grid")

	this.Data["GRId"] = grid

	// 获取该毕业要求的所有指标点
	ips, _ := models.GetMajorGRAllIndicatorPoint(grid)
	logs.Info("指标点：", ips)
	this.Data["Ips"] = ips
}

func (this *MajorGRIPModifyController) Post() {
	this.Ctx.Request.ParseForm()

	reqBody := this.Ctx.Request.Form
	grid := this.Input().Get("grid")
	ips := make([]*common.GRIPType, 0)

	for i := 0; i < len(reqBody["IPNumber"]); i++ {
		ips = append(ips, &common.GRIPType{reqBody["IPNumber"][i], reqBody["IPDescription"][i]})
	}

	// 毕业要求指标点
	err := models.ModifyMajorGRIP(ips, grid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_gr_ip?grid=%s", grid), 302)
}
