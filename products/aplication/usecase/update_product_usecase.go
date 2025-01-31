package aplication

import (
	"holamundo/products/domain/models"
	"holamundo/products/domain/repository"

)

type UpdateProductUsecase struct {
	ProductRepo repository.IProductRepository
}

func (uc *UpdateProductUsecase) Execute (product *models.Product) error{
	return uc.ProductRepo.Update(product)
}