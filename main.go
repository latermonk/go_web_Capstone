package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_web_Capstone/controller"
	"go_web_Capstone/dao"
	"net/http"
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




func main() {
	fmt.Println("Hello , CapStone , I'm Coming !!!")
	//Create database
	//docker run --name mysql8019 --restart=always -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:8.0.19
	//sql: CREATE DATABASE bubble;

	//Connect  to database
	//err := initMySQL()
	err := dao.InitMySQL()

	if err != nil {
		panic(err)
	}
	defer dao.Close()

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

	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	r.GET("/", IndexHandler)

	//TODO
	//v1
	v1Group := r.Group("v1")
	{
		//add POST
		//get the vale ,store to db , show on the final page
		v1Group.POST("/todo", controller.CreateATodo)
		//get
		v1Group.GET("/todo", controller.GetTodoList)
		//v1Group.GET("/todo/:id", func(c *gin.Context) {})
		//PUT
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//DEL
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

	}



	r.Run(":9999")


}
