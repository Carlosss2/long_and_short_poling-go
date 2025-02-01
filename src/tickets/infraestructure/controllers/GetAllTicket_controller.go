package controllers

import (
	"projectC1/src/tickets/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllTicketController struct {
	useCaseGetAll *application.GetAllTicket
}

func NewGetAllTicketController(useCaseGetAll application.GetAllTicket)*GetAllTicketController{
	return &GetAllTicketController{useCaseGetAll: &useCaseGetAll}
}

func (getTicket *GetAllTicketController) View(ctx *gin.Context){
	tickets, err := getTicket.useCaseGetAll.Execute()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
	}
	ctx.JSON(http.StatusOK,gin.H{"tickets":tickets})
}