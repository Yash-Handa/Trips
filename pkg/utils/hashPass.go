package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword creates the hash version of password (string)
func HashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
