package handler

import (
	"net/http"

	"example.com/backend/core/model/aggregate"
	"example.com/backend/core/model/entity"
	"example.com/backend/core/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	var customerHandler = CustomerHandler{}
	customerHandler.customerService = customerService
	return &customerHandler
}

func (ch *CustomerHandler) GetCustomers(c *gin.Context) {
	var customers, _ = ch.customerService.GetAll()
	c.IndentedJSON(http.StatusOK, customers)
}

func (ch *CustomerHandler) GetCustomer(c *gin.Context) {
	// Convert and validate UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	customer, err := ch.customerService.Get(id)

	if err != nil {
		// Handle the error, maybe log it and return an appropriate response
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func (ch *CustomerHandler) SaveCustomer(c *gin.Context) {
	var person entity.Person

	// Bind JSON to customer
	if err := c.BindJSON(&person); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var customer aggregate.Customer
	customer.Person = &person

	// Save the customer
	if err := ch.customerService.Save(customer); err != nil {
		// Handle the error, maybe log it and return an appropriate response
		return
	}

	c.JSON(http.StatusAccepted, person)
}

func (ch *CustomerHandler) UpdateCustomer(c *gin.Context) {
	// Convert and validate UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var person entity.Person

	// Bind JSON to customer
	if err := c.BindJSON(&person); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var customer aggregate.Customer
	customer.Person = &person

	// Update the customer
	if err := ch.customerService.Update(customer, id); err != nil {
		// Handle the error, maybe log it and return an appropriate response
		return
	}

	c.JSON(http.StatusAccepted, customer)
}

func (ch *CustomerHandler) DeleteCustomer(c *gin.Context) {
	// Convert and validate UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Save the customer
	if err := ch.customerService.Delete(id); err != nil {
		// Handle the error, maybe log it and return an appropriate response
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
}
