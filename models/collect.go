package models

import (
  // "fmt"
  _ "log"
  _ "time"
  _ "errors"
  _ "encoding/json"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  _ "github.com/astaxie/beego"
)

/*var (
  db *gorm.DB
  err error
)*/

type CollectLibrary struct {
  gorm.Model
  Mac string
  LibraryGroup string
  Label  string
  Times uint32 `gorm:"default:0"`
};
type CollectSetting struct {
  gorm.Model
  Mac string
  Language string
  SettingGroup string
  Label  string
  Times uint32 `gorm:"default:0"`
};
type CollectJumpOuter struct {
  gorm.Model
  Mac string
  DocsTimes uint32 `gorm:"default:0"`
  ForumTimes uint32 `gorm:"default:0"`
};
type CollectScreen struct{
  gorm.Model
  Mac string
  Pageview string
  Screen string
  Times uint32 `gorm:"default:0"`
};
type CollectCodeLanguage struct{
  gorm.Model
  Mac string
  Language string
  Times uint32 `gorm:"default:0"`
};
type CollectCodeMode struct{
  gorm.Model
  Mac string
  AutoTimes uint32 `gorm:"default:0"`
  CodingTimes uint32 `gorm:"default:0"`
};
/*type CollectError struct{
  gorm.Model
  Mac string
  Group string
  Label string
  Times uint32
};*/
type CollectProduct struct{
  gorm.Model
  Mac string
  ProductName string
  Times uint32 `gorm:"default:0"`
};
type CollectModule struct{
  gorm.Model
  Mac string
  ModuleName string
  Times uint32 `gorm:"default:0"`
};
type CollectPageview struct{
  gorm.Model
  Mac string
  Page string
  Times uint32 `gorm:"default:0"`
};
type collectErrorInfo struct{
  gorm.Model
  Mac string
  Action string
  Label string
  Times uint32 `gorm:"default:0"`
};

func init(){
  /*db, err = gorm.Open("mysql", "root:123456.@/statistic?charset=utf8mb4&parseTime=True&loc=Local")
  db.Debug().DropTableIfExists(&CollectLibrary{},&CollectSetting{},&CollectJumpOuter{},
    &CollectScreen{},&CollectCodeLanguage{}, &CollectCodeMode{},&CollectProduct{},&CollectModule{},&CollectPageview{})
  db.Debug().AutoMigrate(&CollectLibrary{},&CollectSetting{},&CollectJumpOuter{},
    &CollectScreen{},&CollectCodeLanguage{}, &CollectCodeMode{},&CollectProduct{},&CollectModule{},&CollectPageview{})
  fmt.Println(db)*/
}


func AddLibrary(Mac string,  LibraryGroup string, Label string){
  var v CollectLibrary

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("library_group=?", LibraryGroup).Where("label=?", Label).Find(&v)
  // log.Println(v)
  if v.ID == 0{
    v.Mac = Mac
    v.LibraryGroup = LibraryGroup
    v.Label = Label
    v.Times = 1
    // log.Println("insert into")
    db.Create(&v)
  }else{
    v.Times += 1
    db.Save(&v)
  }
  // log.Println(v)
}

func AddSetting(Mac string,  SettingGroup string, Label string){
  var v CollectSetting

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("setting_group=?", SettingGroup).Where("label=?", Label).Find(&v)
  // log.Println(v)
  if v.ID == 0{
    v.Mac = Mac
    v.SettingGroup = SettingGroup
    v.Label = Label
    v.Times = 1
    // log.Println("insert into")
    db.Create(&v)
  }else{
    v.Times += 1
    db.Save(&v)
  }
  // log.Println(v)
}

func AddJumpOuter(Mac string,  Item string){
  var v CollectJumpOuter

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Find(&v)
  // log.Println(v)
  if v.ID == 0{
    v.Mac = Mac
    if(Item == "official documents"){
      v.DocsTimes = 1
    }else if(Item == "online forum"){
      v.ForumTimes = 1
    }
    // log.Println("insert into")
    db.Create(&v)
  }else{
    if(Item == "official documents"){
      v.DocsTimes += 1
    }else if(Item == "online forum"){
      v.ForumTimes += 1
    }
    db.Save(&v)
  }
  // log.Println(v)
}

func AddScreen(Mac string,  Pageview string, Screen string){
  var v CollectScreen

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("pageview=?", Pageview).Where("screen=?", Screen).Find(&v)
  // log.Println(v)
  if v.ID == 0{
    v.Mac = Mac
    v.Pageview = Pageview
    v.Screen = Screen
    v.Times = 1
    // log.Println("insert into")
    db.Create(&v)
  }else{
    v.Times += 1
    db.Save(&v)
  }
  // log.Println(v)
}

func AddCodeMode(Mac string, Codemode string){
  var v CollectCodeMode

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Find(&v)
  // log.Println(v)
  if v.ID == 0{
    v.Mac = Mac
    if(Codemode == "auto"){
      v.AutoTimes = 1
    }else if(Codemode == "coding"){
      v.CodingTimes = 1
    }
    // log.Println("insert into")
    db.Create(&v)
  }else{
    if(Codemode == "auto"){
      v.AutoTimes += 1
    }else if(Codemode == "coding"){
      v.CodingTimes += 1
    }
    db.Save(&v)
  }
  // log.Println(v)
}


func AddCodeLanguage(Mac string, Language string){
  var v CollectCodeLanguage

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("Language=?",Language).Find(&v)
  // log.Println(v)
  if v.ID == 0{
    v.Mac = Mac
    v.Language = Language
    v.Times = 1
    // log.Println("insert into")
    db.Create(&v)
  }else{
    v.Times += 1;
    db.Save(&v)
  }
  // log.Println(v)
}

func AddProduct(Mac string, ProductName string){
  var product CollectProduct

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("product_name=?",ProductName).Find(&product)
  // log.Println(product)
  if product.ID == 0{
    product.Mac = Mac
    product.ProductName = ProductName
    product.Times = 1
    // log.Println("insert into")
    db.Create(&product)
  }else{
    product.Times += 1;
    db.Save(&product)
  }
  // log.Println(product)
}

func AddModule(Mac string, ModuleName string){
  var module CollectModule

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("module_name=?",ModuleName).Find(&module)
  // log.Println(module)
  if module.ID == 0{
    module.Mac = Mac
    module.ModuleName = ModuleName
    module.Times = 1
    // log.Println("insert into")
    db.Create(&module)
  }else{
    module.Times += 1;
    db.Save(&module)
  }
  // log.Println(module)
}

func AddPageView(Mac string, Page string){
  var pageview CollectPageview

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  db.Debug().Where("mac=?", Mac).Where("page=?",Page).Find(&pageview)
  // log.Println(pageview)
  if pageview.ID == 0{
    pageview.Mac = Mac
    pageview.Page = Page
    pageview.Times = 1
    // log.Println("insert into")
    db.Create(&pageview)
  }else{
      pageview.Times += 1;
      db.Save(&pageview)
  }
  // log.Println(pageview)
}


// 收集报错信息
func AddErrorInfo(Mac string, Action string, Label string){
  var error_info collectErrorInfo

  //db.Debug().Where("mac=?", mac).Preload("Cities").Find(&pageview)
  // db.Debug().Where("mac=?", Mac).Where("page=?",Page).Find(&pageview)
  db.Debug().Where(map[string]interface{}{"mac": Mac, "action": Action, "label": Label}).Find(&error_info)
  // db.Debug().map[string]interface{}{"mac": "mac", "page": page, "action": "action", "label": label}.Find(&error_info)
  // log.Println("error_info.ID==", error_info)
  if error_info.ID == 0{
    error_info.Mac = Mac
    error_info.Action = Action
    error_info.Label = Label
    error_info.Times = 1
    // log.Println("insert into")
    db.Create(&error_info)
  }else{
      error_info.Times += 1;
      db.Save(&error_info)
  }
  // log.Println(error_info)
}
