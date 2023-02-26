package query

import (
	"gin-todo/src/db"
	"gin-todo/src/db/models"
)

func GetTodos() (*[]models.Todo, error) {
	var todos []models.Todo
	err := db.DB.Find(&todos).Error

	return &todos, err
}

func GetTodoById(id string) (*models.Todo, error) {
	var todo models.Todo
	err := db.DB.Where("id = ?", id).First(&todo).Error

	return &todo, err
}

func CreateTodo(title string, completed bool) (*models.Todo, error) {
	todo := models.Todo{
		Title:     title,
		Completed: completed,
	}
	err := db.DB.Create(&todo).Error

	return &todo, err
}

func DeleteTodo(id string) error {
	var todo models.Todo
	err := db.DB.Where("id = ?", id).Delete(&todo).Error

	return err
}

func UpdateTodo(id string, title string, competed bool) (*models.Todo, error) {
	var todo, err = GetTodoById(id)
	if err != nil {
		return todo, err
	}

	updateTodoValue := models.Todo{Title: title, Completed: competed}
	err = db.DB.Model(&todo).Updates(updateTodoValue).Error

	return todo, err
}
