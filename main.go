package main

import (
	"holamundo/aplication/usecase"
	"holamundo/core"
	"holamundo/domain/repository"
	"holamundo/infraestructure/controllers"
	"holamundo/infraestructure/routes"
	"log"
	"github.com/gin-gonic/gin" 
)

func main() {
	app := gin.Default()

	core.InitializeApp() //De esta manera comienza el proceso de encendido del servidor

	productRepo := &repository.ProductRepositoryImpl{DB: core.GetDB()} 	// Coneccion del repositorio con la db

	createUC := &aplication.CreateProductUseCase{ProductRepo: productRepo}
	getUC := &aplication.GetProductUseCase{ProductRepo: productRepo}
	updateUC := &aplication.UpdateProductUsecase{ProductRepo: productRepo}
	deleteUC := &aplication.DeleteProductUseCase{ProductRepo: productRepo}
	getAllUC := &aplication.GetAllProductsUseCase{ProductRepo: productRepo}


	createController := &controllers.ProductController{CreateProductUC: createUC}
	getController := &controllers.ProductGetController{GetProductUC: getUC}
	updateContoller := &controllers.ProductUpdateController{UpdateProductUC: updateUC}
	deleteController := &controllers.ProductDeleteController{DeleteProductUC: deleteUC}
	getAllController := &controllers.ProductGetAllController{GetAllProductsUC: getAllUC}
	routes.ProductRoutes(app, createController, getController, updateContoller, deleteController, getAllController)


	log.Println("API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Error al correr el servidor: %v", err)
	}
}
