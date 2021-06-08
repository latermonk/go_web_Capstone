package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_Capstone/models"
	"net/http"
)


func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}


func CreateTodo(c *gin.Context) {
	//fmt.Println("POST")
	var todo models.Todo
	c.BindJSON(&todo)

	err:= models.CreateATodo(&todo)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func GetTodoList(c *gin.Context) {
	//get todo list data
	//var todoList []Todoo

	todoList, err := models.GetAllTodoList()
	if err!= nil{
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

	todo, err := models.GetATodo(id)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)

	//UPDATE DB
	if err = models.UpdateATodo(todo) ;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id);err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}