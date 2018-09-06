package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, password string) (error) {
	byteHashedPassword := []byte(hashedPassword)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return err
	}
	return nil
}