package main

import (
	"go-url-shortner/internal/modules"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	rdb := NewRedisClient("localhost:6379", 0)

	router := gin.Default()
	modules.RegisterRoutes(router, rdb)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
