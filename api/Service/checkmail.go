package service

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func CheckMail(email string) error {
	if err := checkmail.ValidateFormat(email); err != nil {
		return fmt.Errorf("Formato de e-mail inválido: %s", err)
	}
	return nil
}
