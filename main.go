package main

import (
	"go-url-shortner/internal/modules"
	"go-url-shortner/internal/modules/shortener"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	rdb := NewRedisClient("localhost:6379", 0)
	shortenerService := shortener.NewService(rdb)

	router := gin.Default()
	modules.RegisterRoutes(router, shortenerService)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
