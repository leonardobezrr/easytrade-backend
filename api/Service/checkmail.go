package service

import (
	"database/sql"
	"fmt"

	"github.com/badoux/checkmail"
)

func CheckMail(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return fmt.Errorf("Formato de e-mail inválido: %s", err)
	}
	return nil
}

func CheckEmailExists(conn *sql.DB, email string) error {
	var count int
	err := conn.QueryRow("SELECT COUNT(*) FROM usuarios WHERE email = $1", email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("E-mail já está cadastrado!")
	}
	return nil
}
