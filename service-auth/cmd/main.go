package main

import (
	"fmt"
	"log"
	"net/http"

	// Verify env vars or load defaults for local dev

	"github.com/gorilla/mux"
	"github.com/illuminati1911/technews/service-auth/auth/repository"
	"github.com/illuminati1911/technews/service-auth/auth/service"
	"github.com/illuminati1911/technews/utils"
	_ "github.com/illuminati1911/technews/utils/env"
)

func main() {
	db, err := utils.OpenDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repository.NewPSQLAuthRepository(db)
	service.NewAuthService(repo)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":80", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from auth")
}
