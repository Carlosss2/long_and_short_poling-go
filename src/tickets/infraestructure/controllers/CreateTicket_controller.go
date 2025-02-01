package controllers

import (
	"projectC1/src/tickets/application"
	"projectC1/src/tickets/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateTicketController struct {
	useCaseCreate *application.CreateTicket
}

func NewCreateTicketController(useCaseCreate *application.CreateTicket) *CreateTicketController{
	return &CreateTicketController{useCaseCreate: useCaseCreate}
}

func (createTicket *CreateTicketController) Create(c *gin.Context){
	var ticket domain.Ticket

	
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := createTicket.useCaseCreate.Execute(ticket.Client,ticket.Total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		
	}
	c.JSON(http.StatusCreated, gin.H{"message": "ticket registrado"})
}