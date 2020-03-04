package controllers
// package controllers
	// gorm.Model

import (
	// "fmt"
)



var server_config map[string]string

func init() {
	server_config = make(map[string]string)
	server_config = WebConfig()
	// fmt.Println("secret_key2====", GetConfig("secret_key2"))
}

// 生成配置集合
func WebConfig() (map[string]string){
	var server_config map[string]string
	server_config = make(map[string]string)


	server_config["account"] = "17191268"
	server_config["secret_key1"] = "hBKru236Qg2VYm3qTwcNz4br"
	server_config["secret_key2"] = "4W4Zu4CeM7tyABihf5BBN8otzfM3edTS"

	server_config["baiduTokenServer"] = "https://aip.baidubce.com/oauth/2.0/token"
	server_config["word"] = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	server_config["qrcode"] = "https://aip.baidubce.com/rest/2.0/ocr/v1/qrcode"
	server_config["carnumber"] = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"
	server_config["handwriting"] = "https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting"
	server_config["number"] = "https://aip.baidubce.com/rest/2.0/ocr/v1/numbers"

	server_config["face"] = "https://aip.baidubce.com/rest/2.0/face/v3/detect"
	server_config["faceContrast"] = "https://aip.baidubce.com/rest/2.0/face/v3/match"
	server_config["faceMerge"] = "https://aip.baidubce.com/rest/2.0/face/v1/merge"

	server_config["global"] = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"
	server_config["maininfo"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/object_detect"
	server_config["animal"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/animal"
	server_config["plant"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant"
	server_config["flower"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/flower"
	server_config["currency"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/currency"

	server_config["body"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_analysis"
	server_config["body_seg"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_seg"
	server_config["gesture"] = "https://aip.baidubce.com/rest/2.0/image-classify/v1/gesture"


	server_config["speechToText"] = "http://vop.baidu.com/server_api"


	/*server_config["weather"] = "http://www.weather.com.cn/data/cityinfo/"*/
	server_config["weather"] = "https://www.tianqiapi.com/api/"

	/*account string
	secret_key1 string
	secret_key2 string

	baiduTokenServer string
	word string
	qrcode string
	carnumber string
	handwriting string
	number string

	face string
	faceContrast string
	faceMerge string

	global string
	maininfo string
	picture string
	animal string
	plant string
	flower string
	currency string

	body string
	body_seg string
	gesture string


	speechToText string*/

	return server_config
}

// 获取参数
func GetConfig(key string) (string){
	return server_config[key]
}
