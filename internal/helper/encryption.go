package helper

import "golang.org/x/crypto/bcrypt"

func CompareHashedPassword(h, p []byte) bool {
	hashedPassword, password := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	}
	return err == nil
}
