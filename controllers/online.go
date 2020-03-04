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




// tinywebbdb
func (c *OnlineController) Tinywebdb() {
// http://localhost:7077/online/tinywebdb
	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := c.Input().Get("db")
	user := c.Input().Get("user")
	secret := c.Input().Get("secret")
	action := c.Input().Get("action")

	tag := c.Input().Get("tag")
	no := c.Input().Get("no")
	count := c.Input().Get("count")
	value := c.Input().Get("value")
	type_c := c.Input().Get("type")

	fmt.Println(url)

	req_twice := httplib.Get(url)
	req_twice.Param("user", user)
	req_twice.Param("secret", secret)
	req_twice.Param("action", action)
	req_twice.Param("tag", tag)
	req_twice.Param("no", no)
	req_twice.Param("count", count)
	req_twice.Param("value", value)
	req_twice.Param("type", type_c)

	str, err := req_twice.String()
	if err != nil {
	    // t.Fatal(err)
	  // 设置响应状态码
	  c.Ctx.ResponseWriter.WriteHeader(500)
		fmt.Println("tinywebdb error== ",err)
		c.Ctx.WriteString("{error: true}")
		return
	}

	c.Ctx.WriteString(str)
	// 获取前端传递的参数
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

/*图像AI相关*/
// 人脸
func (c *OnlineController) FaceRecognize() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("face")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	fmt.Println("FaceRecognize====")
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) FaceMatch() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("faceContrast")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

// 人体
func (c *OnlineController) BodyAnalysis() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("body")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Gesture() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("gesture")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}


// 文字相关
func (c *OnlineController) GeneralBasic() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("word")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Carnumber() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("carnumber")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Handwriting() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("handwriting")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Numbers() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("number")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

// 图像相关
func (c *OnlineController) AdvancedGeneral() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("global")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) ObjectDetect() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("maininfo")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Animal() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("animal")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Plant() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("plant")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("FaceRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}

func (c *OnlineController) Currency() {

	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))

	url := GetConfig("currency")

	bodyParams := string(c.Ctx.Input.RequestBody)
	token := string(c.Input().Get("access_token"))
	// fmt.Println("tokenid=====",token)

	// 读取参数，根据前端参数向服务器发送请求
	url += "?access_token="+token
	req_twice := httplib.Post(url)
	req_twice.Header("Content-Type","application/json")
	req_twice.Body(bodyParams)
	str, err := req_twice.String()
	// fmt.Println("strRes====", str)
	if err != nil {
	    // t.Fatal(err)
		fmt.Println("CurrencyRecognize error==",err)
	}
	c.Ctx.WriteString(str)
}
