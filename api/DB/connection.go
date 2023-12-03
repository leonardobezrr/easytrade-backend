package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "user_easytrade", "1122", "easytrade")

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
