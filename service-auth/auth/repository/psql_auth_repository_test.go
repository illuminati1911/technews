package repository_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/illuminati1911/technews/models"
	"github.com/illuminati1911/technews/service-auth/auth/repository"
)

func getTestUser() models.User {
	return models.User{
		UserID:    1,
		Username:  "tu",
		Pwhash:    "testhash",
		IsAdmin:   false,
		CreatedAt: "2020-03-22T06:33:14Z",
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tu := getTestUser()

	rows := mock.
		NewRows([]string{"member_id", "username", "pwhash", "is_admin", "created_at"}).
		AddRow(tu.UserID, tu.Username, tu.Pwhash, tu.IsAdmin, tu.CreatedAt)
	mock.ExpectQuery("INSERT INTO member").WillReturnRows(rows).RowsWillBeClosed()

	repo := repository.NewPSQLAuthRepository(db)
	ru, err := repo.CreateUser(tu.Username, tu.Pwhash)
	if err != nil {
		log.Fatal(err)
	}
	if ru.Username != tu.Username {
		t.Errorf("Expected usernames to match")
	}
}

func TestGetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tu := getTestUser()

	rows := mock.
		NewRows([]string{"member_id", "username", "pwhash", "is_admin", "created_at"}).
		AddRow(tu.UserID, tu.Username, tu.Pwhash, tu.IsAdmin, tu.CreatedAt)
	mock.ExpectQuery("SELECT").WillReturnRows(rows).RowsWillBeClosed()

	repo := repository.NewPSQLAuthRepository(db)
	ru, err := repo.GetUserByName(tu.Username)
	if err != nil {
		log.Fatal(err)
	}
	if ru.Username != tu.Username {
		t.Errorf("Expected usernames to match")
	}
}
