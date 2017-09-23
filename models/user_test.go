package models_test

import (
	"testing"
	"github.com/AlbertMorenoDEV/go-for-web-dev/models"
)

func TestAuthenticateFailsWithIncorrectPassword(t *testing.T) {
	user := models.NewUser("gopher@golang.org", "password")
	if user.Authenticate("wrong") {
		t.Errorf("Authenticate should return false for incorrect password")
	}
}

func TestAuthenticateFailsWithCorrectPassword(t *testing.T) {
	user := models.NewUser("gopher@golang.org", "password")
	if !user.Authenticate("password") {
		t.Errorf("Authenticate should return true for incorrect password")
	}
}