package controllers

import (
	"time"
	"strconv"
	"mindplus_statistic/models"
	"github.com/astaxie/beego"
)

type AddUserController struct {
	beego.Controller
}

func (c *AddUserController) Post() {
	var createdTime,updatedTime time.Time

	beego.Debug("input=",c.Input())
	mac := c.Input().Get("mac")
	city := c.Input().Get("city")
	ip := c.Input().Get("ip")
	createdTimeS := c.Input().Get("createdTime")
	updatedTimeS := c.Input().Get("updatedTime")
	version := c.Input().Get("version")
    startuptime := -1
	if v := c.Input().Get("startuptime"); v != ""{
		startuptime, _ = strconv.Atoi(v)
	}

	if(createdTimeS == ""){
		createdTime = time.Now();
	}else{
		createdTime,_ = time.ParseInLocation("2006-01-02 15:04:05", c.Input().Get("createdTime"),time.Local)
	}
	if(updatedTimeS == ""){
		updatedTime = time.Now();
	}else{
		updatedTime,_ = time.ParseInLocation("2006-01-02 15:04:05", c.Input().Get("updatedTime"),time.Local)
	}
	/*beego.Debug(mac)
	beego.Debug(city)
	beego.Debug(ip)
	beego.Debug(version)
	beego.Debug(startuptime)
	beego.Debug(createdTime)
	beego.Debug(updatedTime)*/

	models.AddUser(mac,city,ip,version,startuptime,createdTime,updatedTime)
	c.Ctx.WriteString("ok")
}
