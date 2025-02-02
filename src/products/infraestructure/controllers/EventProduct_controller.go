package controllers

import (
	"net/http"
	"projectC1/src/products/infraestructure"
	"sync"

	"github.com/gin-gonic/gin"
)

type ProductPollingController struct {
	repo             *infraestructure.MySQL
	lastProductCount int  // Almacena el último conteo de productos
	mu               sync.Mutex // Protege el acceso a lastProductCount
}

func NewProductPollingController(repo *infraestructure.MySQL) *ProductPollingController {
	return &ProductPollingController{
		repo:             repo,
		lastProductCount: 0, // Inicializa en 0
	}
}

// Short Polling: Consulta periódicamente cambios en los productos
func (c *ProductPollingController) ShortPolling(ctx *gin.Context) {
	// Obtener el conteo actual de productos desde la base de datos
	currentCount, err := c.repo.GetProductCount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo productos"})
		return
	}

	// Bloquear solo la actualización de lastProductCount
	c.mu.Lock()
	defer c.mu.Unlock()

	// Verificar si hay cambios
	if currentCount != c.lastProductCount {
		c.lastProductCount = currentCount // Actualizar el conteo

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hay cambios",
			"data":    currentCount,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No hay cambios",
		})
	}
}
