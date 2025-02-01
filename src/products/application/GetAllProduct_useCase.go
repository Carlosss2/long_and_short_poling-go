package application

import "projectC1/src/products/domain"

type GetAllProduct struct {
	db domain.IProduct
}

func NewGetAllProduct(db domain.IProduct) *GetAllProduct{
	return &GetAllProduct{db: db}
}

func (gp *GetAllProduct) Execute()([]domain.Product,error){
	return gp.db.GetAll()
}