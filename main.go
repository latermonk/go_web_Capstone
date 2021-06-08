package main

import "github.com/gin-gonic/gin"

func main() {
	//Create database
	//docker run --name mysql8019 --restart=always -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -d mysql:8.0.19
	//sql: CREATE DATABASE bubble;

	//Connect  to database


	r := gin.Default()


	r.Run(":9999")


}
