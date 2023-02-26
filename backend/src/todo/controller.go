package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodosHandler(c *gin.Context) {
	var todos, _ = getTodos()

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func getTodoByIdHandler(c *gin.Context) {
	var todo, err = getTodoById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func createTodoHandler(c *gin.Context) {
	var todo, err = createTodo(c.ShouldBindJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func deleteTodoHandler(c *gin.Context) {
	err := deleteTodo(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func updateTodoHandler(c *gin.Context) {
	var todo, err = updateTodo(c.Param("id"), c.ShouldBindJSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func InitTodoController(r *gin.Engine) {
	r.GET("/todo", getTodosHandler)
	r.GET("/todo/:id", getTodoByIdHandler)
	r.POST("/todo", createTodoHandler)
	r.PUT("/todo/:id", updateTodoHandler)
	r.DELETE("/todo/:id", deleteTodoHandler)
}
