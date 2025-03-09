package repository

import (
	"holamundo/users/domain/models"
)

type IUserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}
