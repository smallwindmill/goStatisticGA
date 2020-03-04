package controllers

import (
	"fmt"
	"mindplus_statistic/models"
	"github.com/astaxie/beego"
)

type CollectController struct {
	beego.Controller
}


func (c *CollectController) GetCollect() {
	mac := c.GetString("_mc")
	ctg := c.GetString("_cctg")
	action := c.GetString("_ac")
	label := c.GetString("_cl")
	ct := c.GetString("_ct")
	content := c.GetString("_cc")
	// beego.Info("mac=",mac)
	// beego.Info("ctg=",ctg)
	// beego.Info("action=",action)
	// beego.Info("label=",label)
	// beego.Info("ct=",ct)
	// fmt.Println("my====", ctg,action)
	// beego.Info("content=",content)
	if(ct == "pageview"){
		models.AddPageView(mac, content)
	}else if(action == "select board or kit"){
		models.AddProduct(mac, label)
	}else if(action == "select small module"){
		models.AddModule(mac, label)
	}else if(ctg == "code"){
		if(action == "select code language"){
			models.AddCodeLanguage(mac,label)
		}else if(action == "study code mode"){
			models.AddCodeMode(mac, label)
		}
	}else if(action == "change arduino screen"){
		models.AddScreen(mac, "arduino", label)
	}else if(action == "change scratch screen"){
		models.AddScreen(mac, "scratch", label)
	}else if(ctg == "jump outer"){
		models.AddJumpOuter(mac,label)
	}else if(ctg == "setting"){
		models.AddSetting(mac, action, label)
	}else if(ctg == "library"){
		models.AddLibrary(mac, action,label)
	}else if(ctg == "error"){
		// fmt.Println("mySelf====", ctg,action);
		// if(action == ""){
			fmt.Println("my error====", c.Input());
		// }
		models.AddErrorInfo(mac, action, label)
	}

	// beego.Info(c.Ctx.Input.URI())
	c.Ctx.WriteString("collect")
}

func (c *CollectController) PostCollect() {
	mac := c.Input().Get("_mc")
	ctg := c.Input().Get("_cctg")
	action := c.Input().Get("_ac")
	label := c.Input().Get("_cl")
	ct := c.Input().Get("_ct")
	content := c.Input().Get("_cc")
	// beego.Info("mac=",mac)
	// beego.Info("ctg=",ctg)
	// beego.Info("action=",action)
	// beego.Info("label=",label)
	// beego.Info("ct=",ct)
	// beego.Info("content=",content)
	if(ct == "pageview"){
		models.AddPageView(mac, content)
	}else if(action == "select board or kit"){
		models.AddProduct(mac, label)
	}else if(action == "select small module"){
		models.AddModule(mac, label)
	}else if(ctg == "code"){
		if(action == "select code language"){
			models.AddCodeLanguage(mac,label)
		}else if(action == "study code mode"){
			models.AddCodeMode(mac, label)
		}
	}else if(action == "change arduino screen"){
		models.AddScreen(mac, "arduino", label)
	}else if(action == "change scratch screen"){
		models.AddScreen(mac, "scratch", label)
	}else if(ctg == "jump outer"){
		models.AddJumpOuter(mac,label)
	}else if(ctg == "setting"){
		models.AddSetting(mac, action, label)
	}else if(ctg == "library"){
		models.AddLibrary(mac, action,label)
	}else if(action == "error"){
		fmt.Println("my error====", c.Input());
		models.AddErrorInfo(mac, action, label)
	}

	// beego.Info(c.Ctx.Input.URI())
	c.Ctx.WriteString("collect success")
}

// func (c *CollectController) getCollectt() {



