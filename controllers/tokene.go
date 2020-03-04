package controllers
// package controllers
	// gorm.Model

import (
	"github.com/astaxie/beego/httplib"
	_"path/filepath"
  _"github.com/astaxie/beego/config"

  "time"
	"fmt"
)

var judge_token map[string]interface{}
// judge_token["access_token"] = ""

func main() {
	go func() {
     for {
        fmt.Println("===================reset tokenid=======================")
        // judge_token["access_token"] = ""
        judge_token = make(map[string]interface{})
        // time.Sleep(time.Second * 60*60*24*29)
        time.Sleep(time.Second * 60*60)
     }
  }()
}


func GetServerTokenid (key1, key2 string) (map[string]interface{}){
  // judge_token = make(map[string]string)
  _, exits := judge_token["access_token"]
	// return judge_token
	fmt.Println("judge_token========", judge_token)
	if(exits && key1 == ""){
		fmt.Println("pass====","pass")
		return judge_token
	}else{
		fmt.Println("no pass=====", "judge_token")
		var token_arr map[string]interface{}
		// 读取参数
	  baiduTokenServer := GetConfig("baiduTokenServer")
  	secret_key1 := key1
		secret_key2 := key2
	  if(key1 == ""){
	  	secret_key1 = GetConfig("secret_key1")
			secret_key2 = GetConfig("secret_key2")
	  }

	  baiduTokenServer += ("?grant_type=client_credentials&client_id="+secret_key1+"&client_secret="+secret_key2)
		token_req := httplib.Get(baiduTokenServer)
		/*token_req.Param("grant_type","client_credentials")
		token_req.Param("client_id", secret_key1)
		token_req.Param("client_secret", secret_key2)*/

		token_req.ToJSON(&token_arr)
		// 判断access_token是否存在
		/*_, token_exits := judge_token["access_token"]
		if token_exits {
			judge_token = token_arr
		}else{
		  token_arr["access_token"] = token_arr["refresh_token"]
		}*/
		// fmt.Println("baiduTokenServer===============",token_arr)
		// baiduTokenServerStr, _ := token_req.String()
		if(key1 == "" && key2 == ""){
			judge_token = token_arr
		}
		fmt.Println("result===============",token_arr)
		// fmt.Println("baiduTokenServerStr===============",baiduTokenServerStr)
		// fmt.Println("baiduTokenServerStr===============",baiduTokenServer)

		// refresh_token
		return token_arr
	}
}
