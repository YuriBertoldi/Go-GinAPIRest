package conexaobd

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DataBase *gorm.DB
	err      error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DataBase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //abre a conexao no PG.
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados, mensagem do sistema" + err.Error())
	}

	CriarTabelas()
}
