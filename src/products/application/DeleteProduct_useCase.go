package application

import "projectC1/src/products/domain"

type DeleteProduct struct {
	db domain.IProduct
}

func NewDeleteProduct(db domain.IProduct) *DeleteProduct{
	return &DeleteProduct{db: db}
}

func (uc *DeleteProduct) Execute(id int32)(error){
	return uc.db.DeleteProduct(id)
}