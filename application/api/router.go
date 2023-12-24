package api

import (
	"example.com/backend/application/handler"
	"example.com/backend/core/service"
	"example.com/backend/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	customerRepository := infrastructure.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(*customerService)

	customer := router.Group("v1/customer")
	{
		customer.GET("/", customerHandler.GetCustomers)
		customer.POST("/", customerHandler.SaveCustomer)
	}
}
