package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}



func CreateATodo(c *gin.Context) {
	//fmt.Println("POST")
	var todo Todo
	c.BindJSON(&todo)

	if err = DB.Create(&todo).Error;err!= nil{
		c.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func GetTodoList(c *gin.Context) {
	//get todo list data
	var todoList []Todo
	if err = DB.Find(&todoList).Error;err!= nil{
		c.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todoList)
	}



}



func  UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error":"invalid id"})
		return
	}
	var todo Todo
	if err = DB.Where("id=?", id).First(&todo).Error;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		return
	}
	c.BindJSON(&todo)
	//UPDATE DB
	if err = DB.Save(&todo).Error;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func  DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error":"invalid id"})
		return
	}
	if err = DB.Where("id=?", id).Delete(Todo{}).Error;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}