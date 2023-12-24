package handler

import (
	"net/http"

	"example.com/backend/core/service"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	var customerHandler = CustomerHandler{}
	customerHandler.customerService = customerService
	return &customerHandler
}

func (handler *CustomerHandler) GetCustomers(c *gin.Context) {
	var customers, _ = handler.customerService.GetAll()
	c.IndentedJSON(http.StatusOK, customers)
}
