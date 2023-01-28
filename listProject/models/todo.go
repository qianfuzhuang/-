package models

import (
	"listProject/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title" `
	Status bool   `json:"status"`
}

// Todo 增删改查
// 创建
func CreateAtodo(todo *Todo) (err error) {
	err = dao.DB.Debug().Create(&todo).Error

	return

}
func GetTodolist() (alllist []*Todo, err error) {
	err = dao.DB.Find(&alllist).Error
	if err != nil {
		return nil, err
	} else {
		return
	}

}
func GetAtodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).First(new(Todo)).Error; err != nil {
		return nil, err
	}
	return
}
func Updata(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}
func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}