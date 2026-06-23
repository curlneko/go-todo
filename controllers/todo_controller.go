package controllers

import (
	"gin-todo/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: 1, Title: "Studying Go.", Completed: false},
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	// JSONを構造体に変換。binding タグをチェック。
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// タイトル重複チェック。大文字小文字を無視してタイトルが同じか比較する。
	for _, v := range todos {
		if strings.EqualFold(v.Title, newTodo.Title) {
			// 409 Conflict
			c.JSON(http.StatusConflict, gin.H{
				"error": "Todo title already exists",
			})
			return
		}
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, v := range todos {
		if strconv.Itoa(v.ID) == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Todo not found",
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var updatedTodo models.Todo

	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, v := range todos {
		if strconv.Itoa(v.ID) == id {
			updatedTodo.ID = v.ID
			todos[i] = updatedTodo

			c.JSON(http.StatusOK, updatedTodo)
			return

		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Todo not found",
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for i, v := range todos {
		if strconv.Itoa(v.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "Todo deleted",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Todo not found",
	})
}
