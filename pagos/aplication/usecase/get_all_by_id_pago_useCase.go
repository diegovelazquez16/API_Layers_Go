package usecase

import (
	"holamundo/pagos/domain/models"
	"holamundo/pagos/domain/repository"
)
type GetPagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *GetPagoUseCase) Execute(id uint) (*models.Pago, error) {
	return uc.PagoRepo.GetByID(id)
}