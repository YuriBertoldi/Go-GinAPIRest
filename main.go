package main

import (
	conexaobd "github.com/YuriBertoldi/Go-GinAPIRest/ConexaoBd"
	routes "github.com/YuriBertoldi/Go-GinAPIRest/routes"
)

func main() {
	conexaobd.Connect()
	routes.HandleRequest()
}
