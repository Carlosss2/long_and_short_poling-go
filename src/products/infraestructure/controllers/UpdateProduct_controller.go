package controllers

import (
	"projectC1/src/products/application"
	"projectC1/src/products/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	useCaseUpdate *application.UpdateProduct
}


func NewUpdateProductController(useCaseUpdate *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCaseUpdate: useCaseUpdate}
}

func (updateProduct *UpdateProductController) Update(c *gin.Context) {
	// Obtener el id del producto de los parámetros de la URL
	id := c.Param("id")

	var product domain.Product

	// Bind JSON para obtener los nuevos valores de nombre y precio
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convertir el id a int32
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Invalido"})
		return
	}

	// Ejecutar la actualización
	err = updateProduct.useCaseUpdate.Execute(int32(productID), product.Name, product.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado"})
}