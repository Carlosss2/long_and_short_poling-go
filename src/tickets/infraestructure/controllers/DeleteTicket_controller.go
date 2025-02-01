package controllers

import (
	"projectC1/src/tickets/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteTicketController struct {
	useCaseDelete *application.DeleteTicket
}

func NewDeleteTicketController(useCaseDelete *application.DeleteTicket)*DeleteTicketController{
	return &DeleteTicketController{useCaseDelete: useCaseDelete}
}

func (deleteTicket *DeleteTicketController) Delete(ctx *gin.Context){
	idParam := ctx.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"invalid"})
		return
	}
	err = deleteTicket.useCaseDelete.Execute(int32(id))
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"Ticket Eliminado"})
}