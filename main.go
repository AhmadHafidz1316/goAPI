package main

import (
	"github.com/AhmadHafidz1316/goAPI/config"
	"github.com/AhmadHafidz1316/goAPI/controllers/productcontroller"
	"github.com/AhmadHafidz1316/goAPI/controllers/usercontroller"
	"github.com/AhmadHafidz1316/goAPI/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	// User Route

	r.GET("users", usercontroller.ShowAll)
	r.POST("/api/register", usercontroller.Register)
	r.POST("/api/login", usercontroller.Login)

	// Product Route

	productRoutes := r.Group("/api")
	productRoutes.Use(config.AuthMiddleware())
	{
		productRoutes.GET("products", productcontroller.Index)
		productRoutes.GET("product/:id", productcontroller.Show)
		productRoutes.POST("product", productcontroller.Create)
		productRoutes.PUT("product/:id", productcontroller.Update)
		productRoutes.DELETE("product/:id", productcontroller.Delete)
	}

	

	r.Run()
}