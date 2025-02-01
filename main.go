package main

import (
	product "projectC1/src/products/infraestructure/dependencies"
	routesProduct "projectC1/src/products/infraestructure/routes"
	

	"github.com/gin-gonic/gin"
)

func main() {
	product.Init()
	

	defer product.CloseDB()
	

	r := gin.Default()
	routesProduct.Routes(r)
	
	r.Run()
	
	
}