package controllers

import (
	"net/http"
	"projectC1/src/tickets/application"
	"projectC1/src/tickets/infraestructure"
	"time"

	"github.com/gin-gonic/gin"
)

type TicketPollingController struct {
	repo *infraestructure.MySQL
}

func NewTicketPollingController(repo *infraestructure.MySQL) *TicketPollingController {
	return &TicketPollingController{repo: repo}
}

// Short Polling: Retorna tickets actuales
func (c *TicketPollingController) ShortPolling(ctx *gin.Context) {
	tickets, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo tickets"})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
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
		ctx.JSON(http.StatusNoContent, nil)
	}
}
