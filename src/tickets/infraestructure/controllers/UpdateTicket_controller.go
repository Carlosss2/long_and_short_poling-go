package controllers

import (
	"projectC1/src/tickets/application"
	"projectC1/src/tickets/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateTicketController struct {
	useCaseUpdate *application.UpdateTicket
}

func NewUpdateTicketController(useCaseUpdate *application.UpdateTicket) *UpdateTicketController{
	return &UpdateTicketController{useCaseUpdate: useCaseUpdate}
}

func (updateTicket *UpdateTicketController) Update(ctx *gin.Context){
	id := ctx.Param("id")

	var ticket domain.Ticket

	if err := ctx.ShouldBindJSON(&ticket); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	ticketId, err := strconv.Atoi(id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"id invalido"})
		return
	}

	err = updateTicket.useCaseUpdate.Execute(int32(ticketId),ticket.Client,ticket.Total)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"message:":"Ticket actualizado"})
}