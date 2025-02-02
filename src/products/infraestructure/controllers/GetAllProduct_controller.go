package controllers

import (
	"projectC1/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	useCaseGetAll *application.GetAllProduct
}

func NewGetAllProductController(useCaseGetAll application.GetAllProduct) *GetAllProductController{
	return &GetAllProductController{useCaseGetAll: &useCaseGetAll}
}

func (getProduct *GetAllProductController) View(ctx *gin.Context){
	products,err := getProduct.useCaseGetAll.Execute() 
	if err != nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		
	}
	ctx.JSON(http.StatusOK,gin.H{"products":products})
}