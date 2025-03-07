package usecase

import (
	"holamundo/pagos/domain/repository"
)
type DeletePagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *DeletePagoUseCase) Execute(id uint) error {
	return uc.PagoRepo.Delete(id)
}