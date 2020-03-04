package controllers

import (
  // "time"
  "fmt"
  // "strconv"
  // "mindpluswebserver/models"
  "github.com/astaxie/beego/httplib"
  "github.com/astaxie/beego"

  "encoding/json"
  // "encoding/base64"
)

type OnlineController struct {
  beego.Controller
}


type speechData struct {
  Speech  string `json:"speech"`
  Format  string `json:"format,omitempty"`
  // Buffer  map[string]string
  // Speech  []byte `json:"speech,omitempty"`
  Token  string `json:"token,omitempty"`
  Rate int `json:"rate"`
  Channel int `json:"channel"`
  Len int `json:"len"`
  Cuid string `json:"cuid"`
}


// 获取tokenID
func (c *OnlineController) GetUserServerTokenid() {
  // tokenid := models.GetServerTokenid()
  // fmt.Println("token_id_global======", tokenid)
  // fmt.Println("token_id_global======", token_id_global)
  // fmt.Println("token_id_global---", Token_id)
  // c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
  key1 := c.Input().Get("key1")
  key2 := c.Input().Get("key2")

  c.Data["json"] = GetServerTokenid(key1, key2)
  c.ServeJSON()
  // c.Ctx.WriteJson(token_id)
  // 获取前端传递的参数
}



// 语音转文字
func (c *OnlineController) SpeechToText() {

  // c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

  url := GetConfig("speechToText")

  token_id := c.Input().Get("access_token")
  // token_id := GetServerTokenid(key1, key2)["access_token"]
  // fmt.Println("input=====",c.Ctx.Input.RequestBody)
  fmt.Println("input=====",c.Ctx.Input.RequestBody)
  // fmt.Println("input=====",string(c.Ctx.Input.RequestBody))

  var ob speechData
  var err error

  json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

  // var token_arr map[string]string

  // var msg =[]byte(ob.Speech)
  // ob.Speech = base64.StdEncoding.EncodeToString(ob.Speech)
  // ob.Speech = base64.StdEncoding.EncodeToString(msg)
  // fmt.Println(base64.StdEncoding.DecodeString(ob.Speech))
  ob.Token = token_id
  // fmt.Println("speechtotext token=====", token_id)
  configdata,_ := json.Marshal(ob)

  // 读取参数，根据前端参数向服务器发送请求

  req_twice := httplib.Post(url)
  req_twice.Header("Content-Type","application/json")


  // fmt.Println("sssssss=", string(configdata))
  req_twice.Body(configdata)
  // req_twice.Body(params)

  // access_token
  str, err := req_twice.String()
  // str, err := req_twice.Response()
  // req_twice.ToJSON(&token_arr)
  fmt.Println("str====", str)
  if err != nil {
      // t.Fatal(err)
    fmt.Println("baiduTokenServer error==",err)
  }

  c.Ctx.WriteString(str)
}



// 天气
func (c *OnlineController) Weather() {

  // c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

  city := c.Input().Get("cityid")
  url := GetConfig("weather")

  fmt.Println(url)
  // var token_arr map[string]string
  // 读取参数

  req_twice := httplib.Get(url)
  // 账号和秘钥在http://www.tianqiapi.com/?action=v1 申请获取
  req_twice.Param("appid", "29549952")
  req_twice.Param("appsecret", "2WngPSdn")
  req_twice.Param("cityid", city)
  req_twice.Param("version", "v6")
  req_twice.Param("vue", "1")
  str, err := req_twice.String()

  if err != nil {
      // t.Fatal(err)
    fmt.Println("weather error== ",err)
    // req_twice.Debug("baiduTokenServer error== ",err)
  }

  c.Ctx.WriteString(str)
  // 获取前端传递的参数
}


