package usecase

import (
	"holamundo/pagos/domain/models"
	"holamundo/pagos/domain/repository"
)

type CreatePagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *CreatePagoUseCase) Execute(pago *models.Pago) error {
	return uc.PagoRepo.Create(pago)
}