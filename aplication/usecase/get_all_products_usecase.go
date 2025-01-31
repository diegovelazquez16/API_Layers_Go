package aplication

import (
	"holamundo/domain/models"
	"holamundo/domain/repository"
)

type GetAllProductsUseCase struct {
	ProductRepo repository.IProductRepository
}

func (uc *GetAllProductsUseCase) Execute() ([]models.Product, error) {
	products, err := uc.ProductRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
