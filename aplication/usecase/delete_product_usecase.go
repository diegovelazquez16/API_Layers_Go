package aplication


import (
	"holamundo/domain/repository"
)

type DeleteProductUseCase struct {
	ProductRepo repository.IProductRepository
}

func (uc *DeleteProductUseCase) Execute (id uint) error{
	return uc.ProductRepo.Delete(id)
}