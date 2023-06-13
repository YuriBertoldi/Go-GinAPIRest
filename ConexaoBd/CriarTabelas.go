package conexaobd

import models "github.com/YuriBertoldi/Go-GinAPIRest/Models"

func CriarTabelas() {
	criarTbCliente()
}

func criarTbCliente() {
	DataBase.AutoMigrate(&models.Cliente{})
}
