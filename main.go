package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Default DB
	})

	router := gin.Default()

	router.GET("/:shortURL", redirectToOriginal)

	router.POST("/shorten", shortenURL)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}

func shortenURL(c *gin.Context) {
	originalURL := c.PostForm("url")

	shortURL, err := generateShortURL(originalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
		return
	}

	err = rdb.Set(ctx, shortURL, originalURL, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store short URL in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func generateShortURL(originalURL string) (string, error) {
	// You can implement your custom logic here to generate a short URL
	// For simplicity, this example uses a basic hash function
	hash := fmt.Sprintf("%x", hashCode(originalURL))
	return hash[:8], nil
}

func hashCode(s string) uint32 {
	var h uint32
	for _, c := range s {
		h = 31*h + uint32(c)
	}
	return h
}

func redirectToOriginal(c *gin.Context) {
	shortURL := c.Param("shortURL")

	originalURL, err := rdb.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve original URL"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}
