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
	r.GET("/user/:id",handlers.GetUser);
	r.POST("/user",handlers.CreateUser);

	r.GET("/product:id",handlers.GetProduct);
	r.POST("/products",handlers.CreateProduct);
	// r.GET("/products", handlers.ListProducts);
	r.PUT("/products:id",handlers.UpdateProduct);
	r.DELETE("/products:id",handlers.DeleteProduct);

	r.GET("/image-analysis", handlers.ImageAnalysis);

	r.Run(":8888");
}