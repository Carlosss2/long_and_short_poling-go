package controllers

import (
	"projectC1/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	useCaseDelete *application.DeleteProduct
}

func NewDeleteProductController(useCaseDelete *application.DeleteProduct)*DeleteProductController{
	return &DeleteProductController{useCaseDelete: useCaseDelete}
}

func (deleteProduct *DeleteProductController) DeleteProduct(ctx *gin.Context){
	idParam := ctx.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error":"Invalid"})
		return
	}
	err = deleteProduct.useCaseDelete.Execute(int32(id))
	if err!= nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"El producto se elimino correctamente"})
}