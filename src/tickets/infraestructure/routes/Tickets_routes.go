package routes

import (
	
	"projectC1/src/tickets/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/tickets")
	createTicket := dependencies.GetCreateTicketController().Create
	getAllTickets := dependencies.GetGetAllTicketController().View
	deleteTickets := dependencies.GetDeleteTicketController().Delete
	updateTickets := dependencies.GetUpdateTicketController().Update


	routes.POST("/",createTicket)
	routes.GET("/",getAllTickets)
	routes.DELETE("/:id",deleteTickets)
	routes.PUT("/:id",updateTickets)
}