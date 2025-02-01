package application

import "projectC1/src/products/domain"

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (uc *CreateProduct) Execute(name string, price float32) error {
	return uc.db.Save(name, price)
}