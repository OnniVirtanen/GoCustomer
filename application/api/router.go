package api

import (
	"database/sql"

	"example.com/backend/application/handler"
	"example.com/backend/core/service"
	"example.com/backend/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, db *sql.DB) {

	customerRepository := infrastructure.NewCustomerRepository(db) //infrastructure.NewCustomerMemoryRepository()
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(*customerService)

	customer := router.Group("v1/customer")
	{
		customer.GET("/", customerHandler.GetCustomers)
		customer.GET("/:id", customerHandler.GetCustomer)
		customer.POST("/", customerHandler.SaveCustomer)
		customer.PUT("/:id", customerHandler.UpdateCustomer)
		customer.DELETE("/:id", customerHandler.DeleteCustomer)
	}

	/*
		productRepository := infrastructure.NewProductRepository()
		productService := service.NewProductService(productRepository)
		productHandler := handler.NewProductHandler(*productService)

		product := router.Group("v1/product")
		{
			product.GET("/", productHandler.GetProducts)
			product.GET("/:id", productHandler.getProduct)
			product.POST("/", productHandler.SaveProduct)
			product.PUT("/:id", productHandler.UpdateProduct)
			product.DELETE("/:id", productHandler.DeleteProduct)
		}
	*/
}
