package launch

import (
	"holamundo/core"
	pedidoUsecase "holamundo/pedidos/aplication/usecase"
	pedidoRepo "holamundo/pedidos/domain/repository"
	pedidoControllers "holamundo/pedidos/infraestructure/controllers"
	pedidoRoutes "holamundo/pedidos/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterPedidoModule(router *gin.Engine) {
	pedidoRepo := &pedidoRepo.PedidosRepositoryImpl{DB: core.GetDB()}

	createPedidoUC := &pedidoUsecase.CreatePedidoUseCase{PedidoRepo: pedidoRepo}
	getAllPedidosUC := &pedidoUsecase.GetAllPedidosUseCase{PedidoRepo: pedidoRepo}
	getPedidoUC := &pedidoUsecase.GetPedidosUseCase{PedidoRepo: pedidoRepo}
	updatePedidoUC := &pedidoUsecase.UpdatePedidoUseCase{PedidoRepo: pedidoRepo}
	deletePedidoUC := &pedidoUsecase.DeletePedidoUseCase{PedidoRepo: pedidoRepo}

	pedidoCreateController := &pedidoControllers.PedidoCreateController{CreatePedidoUC: createPedidoUC}
	pedidoGetAllController := &pedidoControllers.PedidoGetAllController{GetAllPedidosUC: getAllPedidosUC}
	pedidoGetController := &pedidoControllers.PedidoGetController{GetPedidoUC: getPedidoUC}
	pedidoUpdateController := &pedidoControllers.PedidoUpdateController{UpdatePedidoUC: updatePedidoUC}
	pedidoDeleteController := &pedidoControllers.PedidoDeleteController{DeletePedidoUC: deletePedidoUC}

	pedidoRoutes.PedidosRoutes(router, pedidoCreateController, pedidoGetAllController, pedidoUpdateController, pedidoDeleteController, pedidoGetController)
}
