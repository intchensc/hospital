package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"hospital/Model"
	_ "hospital/Model"
	"math/rand"
	"net/http"
)

func main() {
	r := gin.Default()
	//连接数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()

	var beds [36]Model.Bed

	db.AutoMigrate(&Model.Bed{})
	db.AutoMigrate(&Model.Patient{})
	for i := 0; i < 36; i++{
		beds[i].ID = uint(i)
		beds[i].Name = RandStr(4)
		beds[i].Tel = RandStr(11)
		beds[i].Disease = RandStr(7)
		db.Create(&beds[i])
	}


	var beds_out [36]Model.Bed
	for i := 0; i < 36; i++{
		db.Find(&beds_out[i], "ID=?",i+1)
	}

	r.Static("/static","./static")
	r.LoadHTMLGlob("tpl/*")
	//r.SetFuncMap(template.FuncMap{
	//	"subTIme":
	//		func(t time.Time) float64 {
	//			sumH := t.Sub(time.Now())
	//			return sumH.Hours()
	//		},
	//
	//})

	r.GET("/index", func(c *gin.Context) {
		var plist []Model.Patient
		db.Find(&plist)
		c.HTML(http.StatusOK,"tpl/index.tmpl",gin.H{
			"Patient":plist,
			"Bed":beds_out,
		})
	})

	r.POST("/index/query", func(c *gin.Context) {
		tel := c.PostForm("tel")
		var pp = []Model.Patient{}
		if tel == ""{
			var plist []Model.Patient
			db.Find(&plist)
			c.HTML(http.StatusOK,"tpl/index.tmpl",gin.H{
				"Patient":plist,
				"Bed":beds_out,
			})
		}else {
			db.Find(&pp, "Tel=?",tel)
			c.HTML(http.StatusOK,"tpl/index.tmpl",gin.H{
				"Patient":pp,
				"Bed":beds_out,
			})
		}
	})

	r.POST("/index/addpatient", func(c *gin.Context) {
		var add Model.Patient
		if err := c.ShouldBind(&add); err == nil {
			db.Create(&add)
			var plist []Model.Patient
			db.Find(&plist)
			c.HTML(http.StatusOK,"tpl/index.tmpl",gin.H{
				"Patient":plist,
				"Bed":beds_out,
			})
		}
	})

	r.Run(":9090")
}


func RandStr(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//func SubTime(t time.Time, nowtime time.Time) float64 {
//	sumH := t.Sub(nowtime)
//	return sumH.Hours()
//}
//
//func GetNow() time.Time {
//	return time.Now()
//}