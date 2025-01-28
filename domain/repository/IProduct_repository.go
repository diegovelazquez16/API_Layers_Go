package repository

import (
	"gorm.io/gorm"
	"holamundo/domain/models"
)


type IProductRepository interface {
	Create(product *models.Product) error // Interfaz para conectar el modelo con el repositorio
}

type ProductRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}

func (r *ProductRepositoryImpl) Create(product *models.Product) error {
	return r.DB.Create(product).Error
}
