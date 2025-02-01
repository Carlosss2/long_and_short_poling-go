package main

import (
	

	ticket "projectC1/src/tickets/infraestructure/dependencies"
	routesTickets "projectC1/src/tickets/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	ticket.Init()

	
	defer ticket.CloseDB()

	r := gin.Default()
	
	routesTickets.Routes(r)
	r.Run()
	
	
}