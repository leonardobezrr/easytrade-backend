// connection.go
package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	// Leitura das variáveis de ambiente
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	// Construção da string de conexão
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)

	// Abertura da conexão
	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	// Verificação da conexão
	err = conn.Ping()

	return conn, err
}
