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

	shortPolling := dependencies.GetTicketPollingController().ShortPolling
	longPolling := dependencies.GetTicketPollingController().LongPolling

	routes.GET("/polling", shortPolling)         // Short Polling
	routes.GET("/long-polling", longPolling)

	routes.POST("/",createTicket)
	routes.GET("/",getAllTickets)
	routes.DELETE("/:id",deleteTickets)
	routes.PUT("/:id",updateTickets)
}