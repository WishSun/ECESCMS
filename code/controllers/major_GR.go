package controllers

import (
	"ECESCMS/code/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorGRController struct {
	beego.Controller
}

func (this *MajorGRController) Get() {
	this.TplName = "major_GR.html"

	mid := this.Input().Get("mid")
	this.Data["Mid"] = mid

	action := this.Input().Get("action")
	if action == "delete" {
		this.Delete()
	}

	grs, err := models.GetMajorAllGraduationRequirement(mid)
	if err != nil {
		logs.Error(err)
	} else {
		this.Data["GRs"] = grs
	}
}

func (this *MajorGRController) Post() {
	action := this.Input().Get("action")

	switch action {
	case "add":
		this.Add()
	case "modify":
		this.Modify()
	}
}

// 添加毕业要求基本信息
func (this *MajorGRController) Add() {
	GRNumber := this.Input().Get("GRNumber")
	GRName := this.Input().Get("GRName")
	GRDescription := this.Input().Get("GRDescription")
	GRMajor_id := this.Input().Get("GRMajorId")

	err := models.AddGraduationRequirement(GRNumber, GRName, GRMajor_id, GRDescription)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_gr?mid=%s", GRMajor_id), 302)
}

// 修改毕业要求基本信息
func (this *MajorGRController) Modify() {
	GRId := this.Input().Get("GRId")
	logs.Info("grid:", GRId)
	GRNumber := this.Input().Get("GRNumber")
	GRName := this.Input().Get("GRName")
	GRDescription := this.Input().Get("GRDescription")
	GRMajor_id := this.Input().Get("GRMajorId")

	err := models.ModifyGraduationRequirement(GRId, GRNumber, GRName, GRDescription)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_gr?mid=%s", GRMajor_id), 302)
}

// 删除毕业要求
func (this *MajorGRController) Delete() {
	grid := this.Input().Get("grid")
	mid := this.Input().Get("mid")

	err := models.DeleteGraduationRequirement(grid)
	if err != nil {
		logs.Error(err)
	}

	this.Redirect(fmt.Sprintf("/admin/major_gr?mid=%s", mid), 302)
}
