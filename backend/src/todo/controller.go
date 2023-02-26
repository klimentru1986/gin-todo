package todo

import (
	"gin-todo/src/db"
	"gin-todo/src/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodos(c *gin.Context) {
	var todos []models.Todo
	db.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func getTodoById(c *gin.Context) {
	var todo models.Todo
	err := db.DB.Where("id = ?", c.Param("id")).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func createTodo(c *gin.Context) {
	var input CreateTodoDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Title:     input.Title,
		Completed: false,
	}
	db.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"data": &todo})
}

func deleteTodo(c *gin.Context) {
	var todo models.Todo
	err := db.DB.Where("id = ?", c.Param("id")).Delete(&todo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func updateTodo(c *gin.Context) {
	var todo models.Todo
	err := db.DB.Where("id = ?", c.Param("id")).First(&todo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input UpdateTodoDto
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTodoValue := models.Todo{Title: input.Title, Completed: input.Completed}

	db.DB.Model(&todo).Updates(updateTodoValue)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func InitTodoController(r *gin.Engine) {
	r.GET("/todo", getTodos)
	r.GET("/todo/:id", getTodoById)
	r.POST("/todo", createTodo)
	r.PUT("/todo/:id", updateTodo)
	r.DELETE("/todo/:id", deleteTodo)
}
