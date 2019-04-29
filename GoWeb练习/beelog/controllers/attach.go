package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController) Get() {
	// 要将URL中文件名解编码处理，然后去掉第一个字符，因为第一个字符是"/"
	// 例如:   /attachment/temp.go
	filePath := this.Ctx.Request.URL.String()[1:]

	logs.Info("文件路径:", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
	}
	defer f.Close()

	// 将文件内容写到响应中
	_, err = io.Copy(this.Ctx.ResponseWriter, f)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
