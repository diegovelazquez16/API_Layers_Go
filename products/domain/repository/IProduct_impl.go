package repository

import (
	"gorm.io/gorm"
	"holamundo/products/domain/models"
)
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