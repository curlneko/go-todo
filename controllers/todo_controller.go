package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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

	// JSONを構造体に変換。binding タグをチェック。
	if err := c.ShouldBindJSON(&newTodo); err != nil {
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
	id, err := strconv.Atoi(c.Param("id"))

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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var updatedTodo models.Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
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
	id, err := strconv.Atoi(c.Param("id"))

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
