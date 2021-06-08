package routers

import (
	"github.com/gin-gonic/gin"
	"go_web_Capstone/controller"
	"go_web_Capstone/setting"
	//"go_web_Capstone/controller"
)


func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
		r := gin.Default()

		//tel gin framework where is the static
		r.Static("/static", "static")
		r.LoadHTMLGlob("templates/*")

		r.GET("/", controller.IndexHandler)

		//TODO
		//v1
		v1Group := r.Group("v1")
		{
			//add POST
			//get the vale ,store to db , show on the final page
			v1Group.POST("/todo", controller.CreateTodo)
			//get
			v1Group.GET("/todo", controller.GetTodoList)
			//v1Group.GET("/todo/:id", func(c *gin.Context) {})
			//PUT
			v1Group.PUT("/todo/:id", controller.UpdateATodo)
			//DEL
			v1Group.DELETE("/todo/:id", controller.DeleteATodo)

		}

		return r
	}
