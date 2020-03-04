package controllers

import (
	"fmt"
	"os"
	"path"
	"time"
	"strconv"
	"strings"
	"mindplus_statistic/models"
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Post() {
	var feedBackMsg string
	var email string
	var mac string
	var attachmentPath string
	var feedBackUrl string
	var dir string
	beego.Debug("input=",c.Input())

	fmt.Println("file======", c);

	if v := c.Input().Get("feedBackMsg");v != ""{
		feedBackMsg = v
	}
	if v := c.Input().Get("email");v != ""{
		email = v
	}
	beego.Debug("mac==11", c.Input().Get("mac"));
	if v := c.Input().Get("mac");v != ""{
		mac = v
		// dir = "./files/"+strings.Replace(mac, ":", "-", -1 )+"/" //window测试地址
		// attachmentPath = "./files/"+strings.Replace(mac, ":", "-", -1 )+"/"
		dir = "/home/dfrobot/attachment/lib/files/"+strings.Replace(mac, ":", "-", -1 )+"/"
		attachmentPath = "/home/dfrobot/attachment/lib/files/"+strings.Replace(mac, ":", "-", -1 )+"/"
		os.MkdirAll(dir, os.ModePerm)
	}else{
		c.Ctx.WriteString("mac empty")
		return
	}
	attachment := ""

	if(c.Ctx.Request.MultipartForm != nil){
		files := c.Ctx.Request.MultipartForm.File
		for k,v := range files{
			beego.Debug(k)
			beego.Debug(v)
			file, header, e := c.GetFile(k)
			if(e != nil){
				beego.Debug("can not find file  "+k)
			}else{
				file.Close()
				filenameWithSuffix := path.Base(header.Filename)
				fileSuffix := path.Ext(filenameWithSuffix)
				beego.Debug(header.Filename)
				filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
				fn := filenameOnly+strconv.FormatInt(time.Now().Unix(),10)+fileSuffix
				c.SaveToFile(k, dir+fn)
				if(attachment != ""){
					attachment = attachment+","
				}
				attachment = attachment + attachmentPath+fn
			}
		}
	}
  feedBackUrl = strings.Replace(attachment, "/home/dfrobot/attachment/", "", -1 ) //前端读取文件路径
	models.AddFeedBack(mac,time.Now(),feedBackMsg,email,feedBackUrl)

	// if (attachment != "") {
		fmt.Println("attachment====", attachment);
		models.SendMailInit(email, feedBackMsg, attachment) //将用户信息以邮件发送至项目组邮箱
	// }
	//file, header, e := c.GetFile("fileUpload0")
	//file.Close()
	//file, header, e := c.FormFile("fileUpload1")
	//beego.Debug(file)
	//beego.Debug(header)
	//header = header
	//beego.Debug(e)
	beego.Debug(feedBackMsg)
	beego.Debug(email)
	//beego.Debug(strings.Replace(mac, ":", "-", -1 ))
	/*if(url != ""){
		url = "/home/dfrobot/attachment/"+url
		_, err := os.Stat(url)
		if err == nil {
			c.Ctx.Output.Download(url)
		}else{
			c.Ctx.WriteString("file do not exist")
		}
	}*/
	c.Ctx.WriteString("{'status':'ok'}")
}
