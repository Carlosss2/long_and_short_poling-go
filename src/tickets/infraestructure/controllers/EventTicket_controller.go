package controllers

import (
	"net/http"
	"projectC1/src/tickets/application"
	"projectC1/src/tickets/infraestructure"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type TicketPollingController struct {
    repo           *infraestructure.MySQL
    lastTicketCount int // Último conteo de tickets
    mu             sync.Mutex // Protege el acceso a lastTicketCount
}

func NewTicketPollingController(repo *infraestructure.MySQL) *TicketPollingController {
    return &TicketPollingController{
        repo: repo,
        lastTicketCount: 0,
    }
}

func (c *TicketPollingController) ShortPolling(ctx *gin.Context) {
    // Obtener el valor actual de la base de datos
    currentCount, err := c.repo.GetTicketCount()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo tickets"})
        return
    }

    c.mu.Lock() // Bloquear solo la actualización
    defer c.mu.Unlock()

    // Verificar si hay cambios en la cantidad de tickets
    if currentCount != c.lastTicketCount {
        c.lastTicketCount = currentCount // Actualizar el conteo
        
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

// Long Polling: Espera cambios en los tickets
func (c *TicketPollingController) LongPolling(ctx *gin.Context) {
	select {
	case <-application.WaitForTicketUpdate():
		tickets, err := c.repo.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo tickets"})
			return
		}
		ctx.JSON(http.StatusOK, tickets)
	case <-time.After(30 * time.Second): // Timeout de 30 segundos
		ctx.JSON(http.StatusOK, gin.H{"message": "No hay cambios"})
	}
}
