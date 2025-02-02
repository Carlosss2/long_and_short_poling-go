package controllers

import (
	"projectC1/src/products/application"
	"projectC1/src/products/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCaseCreate *application.CreateProduct
	
}//
	func NewCreateProductController(useCaseCreate *application.CreateProduct) *CreateProductController{
		return &CreateProductController{useCaseCreate: useCaseCreate}
	}

	func (createProduct *CreateProductController) Create(c *gin.Context) {
		var product domain.Product
	
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	
		 err := createProduct.useCaseCreate.Execute(product.Name, product.Price)
		 if err != nil {
		 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		 	return
		 }
		 application.NotifyProductUpdate() // Notificar a los listeners
		c.JSON(http.StatusCreated, gin.H{"message": "producto registrado con exito"})
	

	}