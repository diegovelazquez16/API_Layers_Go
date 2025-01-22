package aplication

import "holamundo/domain"

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db:db}
}

func (uc *CreateProduct) Run() {
	uc.db.Save()
}
// Los status no le interesan al negocio solo a los programadores