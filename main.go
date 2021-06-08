package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	DB.AutoMigrate(&Todo{})

	//Create record
	u1 := Todo{001, "tt01", true }

	var t Todo
	db.First(&t)
	//fmt.Print()






	r := gin.Default()


	r.Run(":9999")


}
