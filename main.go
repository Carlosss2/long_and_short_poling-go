package main

import (
	product "projectC1/src/products/infraestructure/dependencies"
	routesProduct "projectC1/src/products/infraestructure/routes"

	ticket "projectC1/src/tickets/infraestructure/dependencies"
	routesTickets "projectC1/src/tickets/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar productos
	product.Init()
	defer product.CloseDB()

	// Inicializar tickets
	ticket.Init()
	defer ticket.CloseDB()

	r := gin.Default()

	// Rutas de productos
	routesProduct.Routes(r)

	// Rutas de tickets
	routesTickets.Routes(r)

	r.Run()
}
