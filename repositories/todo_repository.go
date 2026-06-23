package repositories

import (
	"gin-todo/models"
)

var todos = []models.Todo{
	{ID: 1, Title: "Studying Go.", Completed: false},
}

func GetAll() []models.Todo {
	return todos
}

func Create(todo models.Todo) models.Todo {
	todos = append(todos, todo)
	return todo
}

func Update(id int, todo models.Todo) models.Todo {
	todos[id] = todo
	return todo
}

func Delete(id int) {
	todos = append(todos[:id], todos[id+1:]...)
}
