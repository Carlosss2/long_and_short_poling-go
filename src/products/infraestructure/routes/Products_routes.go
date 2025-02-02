package routes

import (
	
	"projectC1/src/products/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/products")
	createProduct := dependencies.GetCreateProductController().Create
	viewProduct := dependencies.GetAllProductController().View
	
	deleteProduct := dependencies.GetDeleteController().DeleteProduct
	updateProduct := dependencies.GetUpdateController().Update

	longPolling := dependencies.GetProductPollingController().LongPolling

	routes.GET("/long-polling", longPolling) // Long Polling para productos

	routes.POST("/", createProduct)        // POST /products
	routes.GET("/", viewProduct)          // GET /products
	
	routes.DELETE("/:id", deleteProduct)  // DELETE /products/:id
	routes.PUT("/:id", updateProduct)     // PUT /products/:id

}