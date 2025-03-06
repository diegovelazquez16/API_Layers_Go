package bootstrap

import (
	"holamundo/core"
	productUsecase "holamundo/products/aplication/usecase"
	productRepo "holamundo/products/domain/repository"
	productControllers "holamundo/products/infraestructure/controllers"
	productRoutes "holamundo/products/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterProductModule(router *gin.Engine) {
	productRepo := &productRepo.ProductRepositoryImpl{DB: core.GetDB()}

	createProductUC := &productUsecase.CreateProductUseCase{ProductRepo: productRepo}
	getProductUC := &productUsecase.GetProductUseCase{ProductRepo: productRepo}
	updateProductUC := &productUsecase.UpdateProductUsecase{ProductRepo: productRepo}
	deleteProductUC := &productUsecase.DeleteProductUseCase{ProductRepo: productRepo}
	getAllProductsUC := &productUsecase.GetAllProductsUseCase{ProductRepo: productRepo}

	productCreateController := &productControllers.ProductController{CreateProductUC: createProductUC}
	productGetController := &productControllers.ProductGetController{GetProductUC: getProductUC}
	productUpdateController := &productControllers.ProductUpdateController{UpdateProductUC: updateProductUC}
	productDeleteController := &productControllers.ProductDeleteController{DeleteProductUC: deleteProductUC}
	productGetAllController := &productControllers.ProductGetAllController{GetAllProductsUC: getAllProductsUC}

	productRoutes.ProductRoutes(router, productCreateController, productGetController, productUpdateController, productDeleteController, productGetAllController)
}
