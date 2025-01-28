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


	createProductUC := &aplication.CreateProductUseCase{ProductRepo: productRepo}

	productController := &controllers.ProductController{
		CreateProductUC: createProductUC,
	}

	routes.ProductRoutes(app, productController)

	log.Println("API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Error al correr el servidor: %v", err)
	}
}
