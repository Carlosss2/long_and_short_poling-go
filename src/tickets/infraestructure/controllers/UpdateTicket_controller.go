package controllers

import (
	"net/http"
	"projectC1/src/tickets/application"
	"projectC1/src/tickets/domain"
	"projectC1/src/tickets/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateTicketController struct {
	useCaseUpdate *application.UpdateTicket
	repository    *infraestructure.MySQL
	
}

func NewUpdateTicketController(useCaseUpdate *application.UpdateTicket) *UpdateTicketController {
	return &UpdateTicketController{useCaseUpdate: useCaseUpdate}
}

func (updateTicket *UpdateTicketController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var ticket domain.Ticket

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = updateTicket.useCaseUpdate.Execute(int32(ticketId), ticket.Client, ticket.Total)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Notificar a los clientes de Long Polling sobre la actualización del ticket
	mu.Lock()
	tickets, _ := updateTicket.repository.GetAll() // Obtener todos los tickets
	ticketUpdates <- tickets
	mu.Unlock()

	ctx.JSON(http.StatusOK, gin.H{"message": "Ticket actualizado"})
}
