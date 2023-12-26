package handler

import (
	"net/http"

	"example.com/backend/core/model/entity"
	"example.com/backend/core/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	var productHandler = ProductHandler{}
	productHandler.productService = productService
	return &productHandler
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	var products, _ = h.productService.GetAll()
	c.IndentedJSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	// Convert and validate UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product, err := h.productService.Get(id)

	if err != nil {
		// Handle the error, maybe log it and return an appropriate response
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}

func (h *ProductHandler) SaveProduct(c *gin.Context) {
	var product entity.Product

	// Bind JSON to product
	if err := c.BindJSON(&product); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	savedProduct, err := h.productService.Save(product)
	if err != nil {
		// Handle the error, maybe log it and return an appropriate response
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, savedProduct)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	// Convert and validate UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var product entity.Product

	// Bind JSON to product
	if err := c.BindJSON(&product); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Update the product
	if err := h.productService.Update(product, id); err != nil {
		// Handle the error, maybe log it and return an appropriate response
		return
	}

	c.JSON(http.StatusAccepted, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	// Convert and validate UUID
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Save the product
	if err := h.productService.Delete(id); err != nil {
		// Handle the error, maybe log it and return an appropriate response
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
}
