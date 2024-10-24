package model

import (
	"Todolist/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo Todo) (err error) {
	if err := dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodolist() (todos []Todo, err error) {
	if err := dao.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

func UpdateATodo(todo *Todo) (err error) {
	if err = dao.DB.Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(todo *Todo) (err error) {
	if err = dao.DB.Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
