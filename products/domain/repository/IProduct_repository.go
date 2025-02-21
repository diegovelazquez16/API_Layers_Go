package repository

import (
	"holamundo/products/domain/models"
)


type IProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error) 
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}


