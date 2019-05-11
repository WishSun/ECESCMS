package controllers

import (
	"ECESCMS/code/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MajorViewController struct {
	beego.Controller
}

func (this *MajorViewController) Get() {
	this.TplName = "major_view.html"
	mid := this.Input().Get("mid")

	// 获取专业基本信息
	major, _ := models.GetMajorBase(mid)
	this.Data["Major"] = major

	// 获取培养目标信息
	trts, _, _ := models.GetMajorAllTrainTarget(mid)
	this.Data["TrTs"] = trts

	// 获取毕业要求信息
	grs, _ := models.GetMajorAllGraduationRequirement(mid)
	this.Data["GRs"] = grs
}

// 处理网页端动态获取毕业要求的指标点信息
func (this *MajorViewController) Post() {
	// 获取请求携带的grid，然后获取该毕业要求的所有指标点
	ips, err := models.GetMajorGRAllIndicatorPoint(this.Input().Get("grid"))
	if err != nil {
		logs.Error(err)
	}

	// 获取表单数据 构造一个返回对象
	m := &ips
	// 序列化为json后返回
	data, _ := json.Marshal(m)
	fmt.Fprintln(this.Ctx.ResponseWriter, string(data))
}
