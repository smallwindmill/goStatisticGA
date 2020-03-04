package routers

import (
	"mindplus_statistic/controllers"
	// "mindplus_statistic/plugin"

	"github.com/astaxie/beego"
)


func init() {

	// 跨域
	/*beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
      AllowAllOrigins:  true,
      //AllowOrigins:      []string{"https://192.168.0.102"},
      AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
      AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
      ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
      AllowCredentials: true,
  }))*/

	beego.Router("/", &controllers.MainController{})
	beego.Router("/statistic", &controllers.StatisticController{})
	beego.Router("/adduser", &controllers.AddUserController{})
	beego.Router("/getusers", &controllers.GetUsersController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/addfeedback", &controllers.AddFeedBackController{})
	beego.Router("/getfeedbacks", &controllers.GetFeedBacksController{})
	beego.Router("/download", &controllers.DownloadController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/collect", &controllers.CollectController{}, "get:GetCollect;post:PostCollect")


	// 在线http请求转发
	beego.Router("/online/getServerTokenid", &controllers.OnlineController{}, "get:GetUserServerTokenid")
	beego.Router("/online/speechToText", &controllers.OnlineController{}, "get:SpeechToText;post:SpeechToText")
	beego.Router("/online/tinywebdb", &controllers.OnlineController{}, "get:Tinywebdb")
	beego.Router("/online/weather", &controllers.OnlineController{}, "get:Weather")
	beego.Router("/online/face/v3/detect", &controllers.OnlineController{}, "post:FaceRecognize")
	beego.Router("/online/face/v3/match", &controllers.OnlineController{}, "post:FaceMatch")

	beego.Router("/online/ocr/v1/general_basic", &controllers.OnlineController{}, "post:GeneralBasic")
	beego.Router("/online/ocr/v1/license_plate", &controllers.OnlineController{}, "post:Carnumber")
	beego.Router("/online/ocr/v1/carnumber", &controllers.OnlineController{}, "post:Carnumber")
	beego.Router("/online/ocr/v1/handwriting", &controllers.OnlineController{}, "post:Handwriting")
	beego.Router("/online/ocr/v1/numbers", &controllers.OnlineController{}, "post:Numbers")

	beego.Router("/online/image-classify/v1/body_analysis", &controllers.OnlineController{}, "post:BodyAnalysis")
	beego.Router("/online/image-classify/v1/gesture", &controllers.OnlineController{}, "post:Gesture")

	beego.Router("/online/image-classify/v2/advanced_general", &controllers.OnlineController{}, "post:AdvancedGeneral")
	beego.Router("/online/image-classify/v1/object_detect", &controllers.OnlineController{}, "post:ObjectDetect")
	beego.Router("/online/image-classify/v1/animal", &controllers.OnlineController{}, "post:Animal")
	beego.Router("/online/image-classify/v1/plant", &controllers.OnlineController{}, "post:Plant")
	beego.Router("/online/image-classify/v1/currency", &controllers.OnlineController{}, "post:Currency")

}


