package main

import (
	"project-management-application/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();

	r.GET("/user",handlers.GetUser);
	r.POST("/user",handlers.CreateUser);

	r.GET("/product",handlers.GetProduct);
	r.POST("/product",handlers.CreateProduct);

	r.Run(":8888");
}