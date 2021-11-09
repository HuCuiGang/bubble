package models

import "github.com/HuCuiGang/bubble/dao"

/*
	模型层
 */
type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}


func InitModle()  {
	//模型绑定，模型绑定，自动建表
	dao.DB.AutoMigrate(&Todo{})
}

// Todo 增删改查
func CreateATodo(todo *Todo) (err error)  {
	err = dao.DB.Create(&todo).Error
	return
}
func GetAllTodoList() (todoList []*Todo ,err error)  {
	if err = dao.DB.Find(&todoList).Error ;err != nil{
		return nil,err
	}
	return
}

func GetATodo(id string)(todo *Todo ,err error){
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error;err != nil{
		return nil,err
	}
	return
}

func UpdateATodo(todo *Todo )(err error)  {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(id string)(err error){
	err = dao.DB.Where(id).Delete(&Todo{}).Error
	return
}
