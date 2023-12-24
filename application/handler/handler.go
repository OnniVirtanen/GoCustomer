package handler

import (
	"net/http"

	"example.com/backend/core/model/aggregate"
	"example.com/backend/core/model/entity"
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

func (handler *CustomerHandler) SaveCustomer(c *gin.Context) {
	var person entity.Person

	// Bind JSON to customer
	if err := c.BindJSON(&person); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var customer aggregate.Customer
	customer.Person = &person

	// Save the customer
	if err := handler.customerService.Save(customer); err != nil {
		// Handle the error, maybe log it and return an appropriate response
		return
	}

	c.JSON(http.StatusAccepted, person)
}
