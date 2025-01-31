package repository

import (
	"gorm.io/gorm"
	"holamundo/domain/models"
)


type IProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error) 
	GetAll() ([]models.Product, error)
	Update(product *models.Product) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
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
func (r *ProductRepositoryImpl) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
func (r *ProductRepositoryImpl) Update(product *models.Product) error{
	return r.DB.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}
