package usecase

import (
	"holamundo/pagos/domain/models"
	"holamundo/pagos/domain/repository"
)

type UpdatePagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *UpdatePagoUseCase) Execute(pago *models.Pago) error {
	return uc.PagoRepo.Update(pago)
}