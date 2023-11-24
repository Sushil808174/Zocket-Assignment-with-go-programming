package main

import (
	"project-management-application/db"
	"project-management-application/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();

	db.ConnectDB();
	r.GET("/hello",handlers.Hello);
	r.POST("/user",handlers.CreateUserHandler);
	r.POST("/product",handlers.CreateProductHandler);

	r.Run(":8888");
}