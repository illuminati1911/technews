package service_test

import (
	"testing"

	"github.com/illuminati1911/technews/service-auth/auth/repository/mock"
	"github.com/illuminati1911/technews/service-auth/auth/service"
)

func TestCreateUser(t *testing.T) {
	repo := mock.NewMockAuthRepository()
	serv := service.NewAuthService(repo)

	user, err := serv.CreateUser("testUser1", "testPassword1")
	if err != nil {
		t.Errorf("Expected creating user succeed")
	}
	if user.Username != "testUser1" {
		t.Log(user)
		t.Errorf("Expected usernames to match %s", user.Username)
	}
}

func TestLoginSuccess(t *testing.T) {
	repo := mock.NewMockAuthRepository()
	serv := service.NewAuthService(repo)

	_, err := serv.CreateUser("testUser1", "testPassword1")
	if err != nil {
		t.Errorf("Expected creating user succeed")
	}
	token, err := serv.Login("testUser1", "testPassword1")
	if err != nil {
		t.Errorf("Expected login to succeed")
	}
	if len(token) == 0 {
		t.Errorf("Expected to return a token")
	}
}

func TestLoginFailure(t *testing.T) {
	repo := mock.NewMockAuthRepository()
	serv := service.NewAuthService(repo)

	_, err := serv.CreateUser("testUser1", "testPassword1")
	if err != nil {
		t.Errorf("Expected creating user succeed")
	}
	token, err := serv.Login("testUser1", "wrongPassword")
	if err == nil {
		t.Errorf("Expected login to fail")
	}
	if len(token) > 0 {
		t.Errorf("Expected to not return a token")
	}
}
