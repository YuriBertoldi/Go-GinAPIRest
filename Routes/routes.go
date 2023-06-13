package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/YuriBertoldi/Go-GinAPIRest/Controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/teste", controller.Teste)
	r.GET("/clientes", controller.ExibeTodosClientes)
	r.GET("/clientes/:id", controller.ExibeCliente)
	r.GET("/clientes/nome/:nome", controller.ExibeClientePeloNome)
	r.POST("/clientes", controller.NovoCliente)
	r.PUT("/clientes/:id", controller.EditarCliente)
	r.DELETE("/clientes/:id", controller.DeletarCliente)
	r.Run(":5000")
}
