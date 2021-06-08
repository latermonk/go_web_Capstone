package main

import (
	"fmt"
	"go_web_Capstone/dao"
	"go_web_Capstone/models"
	"go_web_Capstone/routers"
	"go_web_Capstone/setting"
)


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
	err := dao.InitMySQL(setting.Conf.MySQLConfig)

	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//bangding MODEL
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run(":9999")


}
