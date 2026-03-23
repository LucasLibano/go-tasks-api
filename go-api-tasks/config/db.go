package config

import (
	"database/sql" //Pacote banco de dados
	"fmt"          //Formatação
	"log"          //Registrar erros
	"os"           //variaveis de ambiente

	"github.com/lucaslibano/go-api-tasks/models"
	_ "github.com/mattn/go-sqlite3" //Driver do SQLite
)

//Conectaao banco de dados

func ConnectDB() *sql.DB {
	dbPatch := os.Getenv("DB_PATH") //Faz a leitura da variavel para o ambiente
	if dbPatch == "" {
		dbPatch = "/data/app.db" //Se não tiver variavel, vai para esse caminho padrao.
	}

	//Abre conexao ao banco

	db, err := sql.Open("sqlite3", dbPatch)
	if err != nil {
		log.Fatal(err)
	}

	//Testa conexao

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	//Cria tabela se não existir

	_, err = db.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database successfully", dbPatch)
	return db
}
