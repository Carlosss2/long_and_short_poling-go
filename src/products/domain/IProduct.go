package domain

type IProduct interface{
	Save(name string, price float32) error
	GetAll()([]Product,error)
	GetById(id int32)(Product,error)
	DeleteProduct(id int32)error
	UpdateProduct(id int32, name string, price float32) error
}