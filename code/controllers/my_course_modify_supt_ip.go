package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MyCourseModifySuptIPController struct {
	beego.Controller
}

// 前台毕业要求更改时需要的后台数据结构
type GRChangeType struct {
	Code int
	GRinfo  *models.GraduationRequirement
	Ips []models.IndicatorPoint
}

// 前台指标点更改时需要的后台数据结构
type IPChangeType struct{
	Code int
	Ip *models.IndicatorPoint
}

func (this *MyCourseModifySuptIPController) Get() {
	this.TplName = "my_course_modify_supt_ip.html"
	majorName := this.Input().Get("major_name")
	this.Data["TeacherName"] = this.GetSession("teacher_name")
	this.Data["MajorName"] = majorName
	this.Data["TSCId"] = this.Input().Get("tscid")
	this.Data["TeTId"] = this.Input().Get("tetid")

	mid, _ := models.GetMajorIdByMajorName(majorName)
	grs, err := models.GetMajorAllGraduationRequirement(mid)
	if err != nil{
		logs.Error(err)
	}
	this.Data["GRs"] = grs

	suptIps, err := models.GetTSCTeachTargetIP(this.Input().Get("tetid"))
	this.Data["SuptIps"] = suptIps
}

func(this *MyCourseModifySuptIPController)Post(){
	this.Ctx.Request.ParseForm()
	tetId := this.Input().Get("tetId")
	tscId := this.Input().Get("tscId")
	majorName := this.Input().Get("majorName")
	ipids := this.Ctx.Request.Form["IPIds"]
	err := models.ModifyTeachtargetSuptIP(tetId, ipids)
	if err != nil{
		logs.Error(err)
	}

	this.Redirect("/teacher/my_course_teach_target?tscid=" + tscId + "&major_name=" + majorName, 302)
}

// 从后台获取毕业要求和指标点信息
func (this *MyCourseModifySuptIPController) GetGRIPs() {
	majorName := this.Input().Get("MajorName")
	grText := this.Input().Get("GRText")

	m := &GRChangeType{}
	gr, ips, err := models.GetGRAndIPs(majorName, grText)
	if err == nil {
		m.Code = 0
		m.GRinfo = gr
		m.Ips = ips
	} else {
		m.Code = 1
	}
	data, _ := json.Marshal(m)
	_, _ = fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}

// 从后台获取指标点信息
func (this *MyCourseModifySuptIPController)GetIP(){
	majorName := this.Input().Get("MajorName")
	grText := this.Input().Get("GRText")
	ipText := this.Input().Get("IPText")

	m := &IPChangeType{}
	ip, err := models.GetIP(majorName, grText, ipText)
	if err == nil{
		m.Code = 0
		m.Ip = ip
	}else{
		m.Code = 1
	}
	data, _ := json.Marshal(m)
	_, _ = fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}