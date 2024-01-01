package api

import (
	"database/sql"

	"example.com/backend/application/handler"
	"example.com/backend/core/service"
	"example.com/backend/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupProductRouter(router *gin.Engine, db *sql.DB) {
	productRepository := infrastructure.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(*productService)

	product := router.Group("v1/product")
	{
		product.GET("/", productHandler.GetProducts)
		product.GET("/:id", productHandler.GetProduct)
		product.POST("/", productHandler.SaveProduct)
		product.PUT("/:id", productHandler.UpdateProduct)
		product.DELETE("/:id", productHandler.DeleteProduct)
	}
}
