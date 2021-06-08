package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type USerInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}


func initMySQL() ( err error){
	dsn := "root:root@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	fmt.Println("Hello , CapStone , I'm Coming !!!")
	//Create database
	//docker run --name mysql8019 --restart=always -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:8.0.19
	//sql: CREATE DATABASE bubble;

	//Connect  to database
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	//bangding MODEL
	DB.AutoMigrate(&Todo{})

	//Create record
	//u1 := Todo{001, "tt01", true }

	//var t Todo
	//db.First(&t)
	//fmt.Print()





	//Create gin route
	r := gin.Default()

	//tel gin framework where is the static
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//TODO
	//v1
	v1Group := r.Group("v1")
	{
		//add POST
		v1Group.POST("/todo", func(c *gin.Context) {
			//fmt.Println("POST")
		})
		//get
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//PUT
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		//DEL
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})

	}



	r.Run(":9999")


}
