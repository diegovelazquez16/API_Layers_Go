package main

import (
	"holamundo/core"
	// Productos
	productUsecase "holamundo/products/aplication/usecase"
	productRepo "holamundo/products/domain/repository"
	productControllers "holamundo/products/infraestructure/controllers"
	productRoutes "holamundo/products/infraestructure/routes"

	// Usuarios
	userUsecase "holamundo/users/aplication/usecase"
	userRepo "holamundo/users/domain/repository"
	userControllers "holamundo/users/infraestructure/controllers"
	userRoutes "holamundo/users/infraestructure/routes"
	"log"


	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	core.InitializeApp() //De esta manera comienza el proceso de encendido del servidor

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

	productRoutes.ProductRoutes(app, productCreateController, productGetController, productUpdateController, productDeleteController, productGetAllController)

	// Usuarios
	userRepo := &userRepo.UserRepositoryImpl{DB: core.GetDB()}

	createUserUC := &userUsecase.CreateUserUseCase{UserRepo: userRepo}
	getAllUsersUC := &userUsecase.GetAllUsersUseCase{UserRepo: userRepo}
	updateUserUC := &userUsecase.UpdateUserUseCase{UserRepo: userRepo}


	userCreateController := &userControllers.UserCreateController{CreateUserUC: createUserUC}
	userGetAllController := &userControllers.UserGetAllController{GetAllUsersUC: getAllUsersUC}
	userUpdateController := &userControllers.UserUpdateController{UpdateUserUC: updateUserUC}


	userRoutes.UserRoutes(app, userCreateController, userGetAllController, userUpdateController)




	log.Println("API corriendo en http://localhost:8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Error al correr el servidor: %v", err)
	}
}
