package services

import (
	"errors"
	"strings"

	"gin-todo/models"
	"gin-todo/repositories"
)

var ErrTodoNotFound = errors.New("Todo not found")
var ErrDuplicateTitle = errors.New("Todo title already exists")

func GetTodos() []models.Todo {
	return repositories.GetAll()
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	todos := repositories.GetAll()

	// タイトル重複チェック。大文字小文字を無視してタイトルが同じか比較する。
	for _, v := range todos {
		if strings.EqualFold(v.Title, todo.Title) {
			return models.Todo{}, ErrDuplicateTitle
		}
	}

	todo.ID = len(todos) + 1

	return repositories.Create(todo), nil
}

func GetTodoByID(id int) (models.Todo, error) {
	todos := repositories.GetAll()

	for _, v := range todos {
		if v.ID == id {
			return v, nil
		}
	}

	return models.Todo{}, ErrTodoNotFound
}

func UpdateTodo(id int, updatedTodo models.Todo) (models.Todo, error) {
	todos := repositories.GetAll()

	for i, v := range todos {

		// 更新対象を探す
		if v.ID == id {

			// 自分以外との重複チェック
			for _, t := range todos {
				if t.ID != id &&
					strings.EqualFold(t.Title, updatedTodo.Title) {

					return models.Todo{}, ErrDuplicateTitle
				}
			}

			updatedTodo.ID = id
			repositories.Update(i, updatedTodo)

			return updatedTodo, nil
		}
	}

	return models.Todo{}, ErrTodoNotFound
}

func DeleteTodo(id int) error {
	todos := repositories.GetAll()

	for i, v := range todos {
		if v.ID == id {
			repositories.Delete(i)
			return nil
		}
	}

	return ErrTodoNotFound
}
