package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-todo/models"
	"gin-todo/responses"
	errResponse "gin-todo/responses/errors"
	"gin-todo/services"
)

func GetTodos(c *gin.Context) {
	todos := services.GetTodos()
	responses.HandleSuccess(c, http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := bindTodoPayload(c, &newTodo); err != nil {
		responses.HandleError(c, err)
		return
	}

	todo, err := services.CreateTodo(newTodo)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleSuccess(c, http.StatusCreated, todo)
}

func GetTodoByID(c *gin.Context) {
	id, err := parseTodoID(c)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	todo, err := services.GetTodoByID(id)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleSuccess(c, http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id, err := parseTodoID(c)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	var updatedTodo models.Todo

	if err := bindTodoPayload(c, &updatedTodo); err != nil {
		responses.HandleError(c, err)
		return
	}

	todo, err := services.UpdateTodo(id, updatedTodo)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleSuccess(c, http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, err := parseTodoID(c)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	err = services.DeleteTodo(id)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleSuccess(c, http.StatusOK, gin.H{
		"message": "Todo deleted",
	})
}

func parseTodoID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, errResponse.ErrInvalidID
	}

	return id, nil
}

func bindTodoPayload(c *gin.Context, target *models.Todo) error {
	if err := c.ShouldBindJSON(target); err != nil {
		return errResponse.ErrInvalidRequest
	}

	return nil
}
