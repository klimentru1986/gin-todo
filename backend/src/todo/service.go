package todo

import (
	"gin-todo/src/db/models"
	"gin-todo/src/db/query"
)

func getTodos() (*[]models.Todo, error) {
	var todos, err = query.GetTodos()

	return todos, err
}

func getTodoById(id string) (*models.Todo, error) {
	var todo, err = query.GetTodoById(id)
	return todo, err
}

func createTodo(bindFunc func(obj any) error) (*models.Todo, error) {
	var input CreateTodoDto
	err := bindFunc(&input)
	if err != nil {
		return nil, err
	}

	var todo, createErr = query.CreateTodo(input.Title, false)
	return todo, createErr
}

func deleteTodo(id string) error {
	return query.DeleteTodo(id)

}

func updateTodo(id string, bindFunc func(obj any) error) (*models.Todo, error) {
	var input UpdateTodoDto
	err := bindFunc(&input)
	if err != nil {
		return nil, err
	}

	var todo, errQuery = query.UpdateTodo(id, input.Title, input.Completed)

	return todo, errQuery
}
