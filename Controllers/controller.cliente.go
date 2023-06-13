package Controllers

import (
	"net/http"

	conexaobd "github.com/YuriBertoldi/Go-GinAPIRest/ConexaoBd"
	models "github.com/YuriBertoldi/Go-GinAPIRest/Models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosClientes(c *gin.Context) {
	var clientes []models.Cliente

	conexaobd.DataBase.Find(&clientes)

	c.JSON(http.StatusOK, &clientes)
}

func ExibeCliente(c *gin.Context) {
	id := c.Params.ByName("id")
	var cliente models.Cliente

	conexaobd.DataBase.First(&cliente, id)

	if cliente.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cliente nao encontrado.",
		})
	} else {
		c.JSON(http.StatusOK, &cliente)
	}
}

func ExibeClientePeloNome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	var cliente models.Cliente

	conexaobd.DataBase.Where(&models.Cliente{Nome: nome}).First(&cliente)

	if cliente.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cliente nao encontrado.",
		})
	} else {
		c.JSON(http.StatusOK, &cliente)
	}
}

func DeletarCliente(c *gin.Context) {
	id := c.Params.ByName("id")
	var cliente models.Cliente

	conexaobd.DataBase.Delete(&cliente, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Cliente deletado com sucesso.",
	})

}

func NovoCliente(c *gin.Context) {
	var cliente models.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	conexaobd.DataBase.Create(&cliente)
	c.JSON(http.StatusCreated, &cliente)
}

func EditarCliente(c *gin.Context) {
	id := c.Params.ByName("id")
	var cliente models.Cliente

	conexaobd.DataBase.First(&cliente, id)

	if cliente.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Cliente nao encontrado.",
		})
		return
	}

	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	conexaobd.DataBase.Model(&cliente).UpdateColumns(cliente)
	c.JSON(http.StatusOK, gin.H{
		"message": "Atualizado com sucesso",
	})
}
