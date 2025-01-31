package repository

import (
	"gorm.io/gorm"
	"holamundo/domain/models"
)


type IProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error) // Interfaz para conectar el modelo con el repositorio
}

type ProductRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}

func (r *ProductRepositoryImpl) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepositoryImpl) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}