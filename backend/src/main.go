package main

import (
	"gin-todo/src/db"
	"gin-todo/src/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	todo.InitTodoController(r)

	r.Run()
}
