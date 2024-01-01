package api

import (
	"database/sql"

	"example.com/backend/application/handler"
	"example.com/backend/core/service"
	"example.com/backend/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupCustomerRouter(router *gin.Engine, db *sql.DB) {
	customerRepository := infrastructure.NewCustomerRepository(db)
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
}
