package aplication

import (
	"holamundo/products/domain/models"
	"holamundo/products/domain/repository"
)

type GetProductUseCase struct {
	ProductRepo repository.IProductRepository
}

func (uc *GetProductUseCase) Execute(id uint) (*models.Product, error) {
	return uc.ProductRepo.GetByID(id)
}
