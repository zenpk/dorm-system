package user

import "golang.org/x/crypto/bcrypt"

func bCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
