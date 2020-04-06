package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illuminati1911/technews/service-auth/auth/delivery/http"
	"github.com/illuminati1911/technews/service-auth/auth/repository"
	"github.com/illuminati1911/technews/service-auth/auth/service"
	"github.com/illuminati1911/technews/utils"

	// Verify env vars or load defaults for local dev
	_ "github.com/illuminati1911/technews/utils/env"
)

func main() {
	db, err := utils.OpenDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repository.NewPSQLAuthRepository(db)
	serv := service.NewAuthService(repo)
	router := gin.Default()
	http.NewAuthHTTPHandler(serv, router)
	router.Run(":80")
}
