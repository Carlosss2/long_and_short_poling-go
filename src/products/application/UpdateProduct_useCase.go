package application

import "projectC1/src/products/domain"

type UpdateProduct struct {
	db domain.IProduct
}

func NewUpdateProduct(db domain.IProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (uc *UpdateProduct) Execute(id int32, name string, price float32) error {
	return uc.db.UpdateProduct(id,name, price)
}