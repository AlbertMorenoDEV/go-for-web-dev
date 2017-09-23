package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `db:"username"`
	Secret   []byte `db:"secret"`
}

func NewUser(username, password string) *User {
	secret, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &User{
		Username: username,
		Secret: secret,
	}
}

func (u *User) Authenticate(password string) bool {
	return bcrypt.CompareHashAndPassword(u.Secret, []byte(password)) == nil
}
