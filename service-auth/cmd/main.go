package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	grpcHandler "github.com/illuminati1911/technews/service-auth/auth/delivery/grpc"
	"github.com/illuminati1911/technews/service-auth/auth/delivery/http"
	"github.com/illuminati1911/technews/service-auth/auth/repository"
	"github.com/illuminati1911/technews/service-auth/auth/service"
	"github.com/illuminati1911/technews/utils"

	// Verify env vars or load defaults for local dev
	_ "github.com/illuminati1911/technews/utils/env"
)

func startGRPC(wg *sync.WaitGroup, grpc *utils.GRPC) {
	if err := grpc.Start(); err != nil {
		panic(err)
	}
	wg.Done()
}

func startHTTP(wg *sync.WaitGroup, router *gin.Engine) {
	router.Run(":80")
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	db, err := utils.OpenDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Initialize servers
	grpc := utils.NewGRPC(":3000")
	router := gin.Default()
	// Repositories, services and delivery handlers
	repo := repository.NewPSQLAuthRepository(db)
	serv := service.NewAuthService(repo)
	grpcHandler.NewAuthGRPCHandler(serv, grpc)
	http.NewAuthHTTPHandler(serv, router)
	// Start servers
	go startHTTP(wg, router)
	go startGRPC(wg, grpc)
	wg.Wait()
}
