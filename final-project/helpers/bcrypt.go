package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	salt := 8
	password := []byte(pass)

	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePassword(h, p []byte) bool {
	hash, password := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
