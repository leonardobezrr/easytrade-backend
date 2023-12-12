package service

import "golang.org/x/crypto/bcrypt"

func HashSenha(senha string) (string, error) {
	hashedSenha, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedSenha), nil
}
