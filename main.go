package main

import (
	"github.com/shogo-nakano-desu/go-todo-app/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
