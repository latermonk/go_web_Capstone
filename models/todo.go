package models

import (
	"go_web_Capstone/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//POST DEL PUT GET

func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	return
}


func GetAllTodoList() (todoList []*Todo, err error){
	//var todoList []*Todo
	if err := dao.DB.Find(&todoList).Error;err != nil{
		return nil, err
	}
	return
}


func GetATodo(id string)(todo *Todo, err error){
	//var todo Todo
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error;err!=nil{
		//c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		return nil, err
	}
	return
}


func UpdateATodo(todo *Todo)(err error){
	err = dao.DB.Create(&todo).Error
	return
}


func DeleteATodo(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}