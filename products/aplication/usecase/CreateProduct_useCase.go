package aplication

import (
	"holamundo/products/domain/models"
	"holamundo/products/domain/repository"
)

type CreateProductUseCase struct {
	ProductRepo repository.IProductRepository
}

func (uc *CreateProductUseCase) Execute(product *models.Product) error {
	return uc.ProductRepo.Create(product)
}



// Los status no le interesan al negocio solo a los programadores