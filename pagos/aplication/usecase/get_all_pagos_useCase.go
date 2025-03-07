package usecase

import (
	"holamundo/pagos/domain/models"
	"holamundo/pagos/domain/repository"
)
type GetAllPagosUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *GetAllPagosUseCase) Execute() ([]models.Pago, error) {
	return uc.PagoRepo.GetAll()
}

