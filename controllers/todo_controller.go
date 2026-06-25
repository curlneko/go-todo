package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	appErr "gin-todo/errors"
	"gin-todo/models"
	"gin-todo/services"
	"gin-todo/utils"
)

func GetTodos(c *gin.Context) {
	todos := services.GetTodos()
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := bindTodoPayload(c, &newTodo); err != nil {
		utils.HandleError(c, err)
		return
	}

	todo, err := services.CreateTodo(newTodo)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func GetTodoByID(c *gin.Context) {
	id, err := parseTodoID(c)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	todo, err := services.GetTodoByID(id)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id, err := parseTodoID(c)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var updatedTodo models.Todo

	if err := bindTodoPayload(c, &updatedTodo); err != nil {
		utils.HandleError(c, err)
		return
	}

	todo, err := services.UpdateTodo(id, updatedTodo)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, err := parseTodoID(c)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	err = services.DeleteTodo(id)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted",
	})
}

func parseTodoID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, appErr.ErrInvalidID
	}

	return id, nil
}

func bindTodoPayload(c *gin.Context, target *models.Todo) error {
	if err := c.ShouldBindJSON(target); err != nil {
		return appErr.ErrInvalidRequest
	}

	return nil
}
