package models

import (
	"fmt"
	_ "log"
	"time"
	"errors"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
var (
	db *gorm.DB
	err error
)

type User struct {
	gorm.Model
	Mac string `gorm:"type:varchar(64);unique_index" json:"mac"`
	City string `json:"city"`
	//Cities []City `gorm:"ForeignKey:CityId" json:"cities"`
	Ip  string `json:"ip"`
	Visits int `gorm:"Default:1" json:"visits"`
	Version string `json:"version"`
	Startup int `json:"startup"`
}
/*
type City struct {
	ID	int `gorm:"primary_key" json:"cid"`
	CityId int `json:"city_id"`
	CityName string `json:"city_name"`
}
*/
type FeedBack struct {
	gorm.Model
	Mac string
	FeedBackMsg string
	Email  string
	Path string
	IsReply bool
};

func AddFeedBack(mac string, created time.Time, feedBackMsg string, email string, path string) error{
	fb := FeedBack{
		Mac:mac,
		FeedBackMsg:feedBackMsg,
		Email:email,
		Path:path,
	}
	db.Debug().Create(&fb)
	return nil
}

func GetFeedBacks(createdStartTime time.Time, createdEndTime time.Time, offset int, limit int) string{
	result := []FeedBack{}
	db.Debug().Model(&FeedBack{}).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Find(&result)
	// log.Println(result)
	data,_ := json.Marshal(result)
	return string(data)

}

func init() {
	return
	// db, err = gorm.Open("mysql", "root:123456.@/statistic?charset=utf8mb4&parseTime=True&loc=Local")
	db, err = gorm.Open("mysql", "root:hidfrobot@/statistic?charset=utf8mb4&parseTime=True&loc=Local")
	db.Debug().AutoMigrate(&User{},&FeedBack{})
	db.Debug().AutoMigrate(&CollectLibrary{},&CollectSetting{},&CollectJumpOuter{},
		&CollectScreen{},&CollectCodeLanguage{}, &CollectCodeMode{},&CollectProduct{},&CollectModule{},&CollectPageview{})
	fmt.Println(db)
	//db.Debug().Model(&City{}).AddForeignKey("city_id", "users(id)", "CASCADE", "CASCADE")

}

func AddUser(mac string , city string , ip string, version string,startuptime int,createdTime time.Time, updatedTime time.Time){
	var user User

	//db.Debug().Where("mac=?", mac).Preload("Cities").Find(&user)
	db.Debug().Where("mac=?", mac).Find(&user)
	// log.Println(user)
	if user.ID == 0{
		user.CreatedAt = createdTime
		user.UpdatedAt = updatedTime
		user.Mac = mac
		user.City = city
		user.Version = version
		user.Startup = startuptime
		//user.Cities = []City{{CityName:city}}
		user.Ip = ip
		// log.Println("insert into")
		db.Create(&user)
	}else{
		// exist := false
		user.Visits += 1;
		user.City = city
		user.Version = version
		user.Startup = startuptime
		// for _,v := range user.Cities{
		// 	if v.CityName == city{
		// 		exist = true
		// 	}
		// }
		//if !exist{
		//	user.Cities = append(user.Cities, City{CityName:city})
		//}
		db.Save(&user)
	}
	// log.Println(user)
}

func GetUser(mac string) string{
	var user User
	//db.Debug().Where("mac=?",mac).Preload("Cities").Find(&user)
	db.Debug().Where("mac=?",mac).Find(&user)
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
	return string(data)
}

func GetUsers(visitsMin int, visitsMax int, updatedStartTime time.Time, updatedEndTime time.Time, createdStartTime time.Time, createdEndTime time.Time, city string, offset int, limit int, orderby string) string {
	users := []User{}
	if(city == ""){
		//db.Offset(offset).Limit(limit).Debug().Preload("Cities").Where("visits > ? AND visits < ?", visitsMin, visitsMax).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Where("updated_at BETWEEN ? AND ?", updatedStartTime, updatedEndTime).Order(orderby).Find(&users)
		db.Offset(offset).Limit(limit).Debug().Where("visits > ? AND visits < ?", visitsMin, visitsMax).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Where("updated_at BETWEEN ? AND ?", updatedStartTime, updatedEndTime).Order(orderby).Find(&users)
	}else{
		//db.Offset(offset).Limit(limit).Debug().Preload("Cities","city_name = ?",city).Where("visits > ? AND visits < ?", visitsMin, visitsMax).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Where("updated_at BETWEEN ? AND ?", updatedStartTime, updatedEndTime).Order(orderby).Find(&users)
		db.Offset(offset).Limit(limit).Debug().Where("visits > ? AND visits < ?", visitsMin, visitsMax).Where("city > ? ", city).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Where("updated_at BETWEEN ? AND ?", updatedStartTime, updatedEndTime).Order(orderby).Find(&users)
	}
	/*N := len(users)
    for i := 0; i < N; i++ {
        for i := 0; i < len(users); i++ {
            if len(users[i].Cities) == 0{
                users = append(users[:i], users[i+1:]...)
                i--
            }
        }
    }*/
	// log.Println(users)
	data,_ := json.Marshal(users)
	return string(data)
}

func GetUserStatistic(visitsMin int, visitsMax int, updatedStartTime time.Time, updatedEndTime time.Time, createdStartTime time.Time, createdEndTime time.Time) string {
	type StatisticEntry struct{
		City string `json:"city"`
		Total int   `json:"total"`
	}
	type Statistic struct{
		Total int `json:"total"`
		Users []StatisticEntry `json:"users"`
	}

	result := Statistic{}

	//db.Debug().Model(&User{}).Preload("Cities").Where("visits > ? AND visits < ?", visitsMin, visitsMax).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Where("updated_at BETWEEN ? AND ?", updatedStartTime, updatedEndTime).Select("city , count(*) as total").Group("city").Scan(&result.Users)
	db.Debug().Model(&User{}).Where("visits > ? AND visits < ?", visitsMin, visitsMax).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Where("updated_at BETWEEN ? AND ?", updatedStartTime, updatedEndTime).Select("city , count(*) as total").Group("city").Scan(&result.Users)
	total := 0
	for _,v := range result.Users{
		total += v.Total
	}
	result.Total = total
	// log.Println(result)
	data,_ := json.Marshal(result)
	return string(data)
}

func GetFeedBackStatistic(createdStartTime time.Time, createdEndTime time.Time) string {
	type FeedBackStatistic struct{
		Total int `json:"total"`
	}

	result := FeedBackStatistic{}

	db.Debug().Model(&FeedBack{}).Where("created_at BETWEEN ? AND ?", createdStartTime, createdEndTime).Select("count(*) as total").Scan(&result)
	// log.Println(result)
	data,_ := json.Marshal(result)
	return string(data)
}

func Login(username string , password string) error{
	if(username == "root") && (password == "jzk0CxcfldM1W8tdJk+IhJFnTtO3Jc0fY0owet57zQs="){
		return nil
	}
	return errors.New("username or password error")
}
