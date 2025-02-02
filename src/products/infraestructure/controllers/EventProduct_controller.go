package controllers

import (
	"net/http"
	"projectC1/src/products/application"
	"projectC1/src/products/infraestructure"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductPollingController struct {
	repo *infraestructure.MySQL
}

func NewProductPollingController(repo *infraestructure.MySQL) *ProductPollingController {
	return &ProductPollingController{repo: repo}
}

// Long Polling: Espera cambios en los productos
func (c *ProductPollingController) LongPolling(ctx *gin.Context) {
	select {
	case <-application.WaitForProductUpdate():
		products, err := c.repo.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo productos"})
			return
		}
		ctx.JSON(http.StatusOK, products)
	case <-time.After(30 * time.Second): // Timeout de 30 segundos
		ctx.JSON(http.StatusNoContent, nil)
	}
}