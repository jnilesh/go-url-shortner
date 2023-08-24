package shortener

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RegisterRoutes(r *gin.Engine, rdb *redis.Client) {
	shortenerGroup := r.Group("/")
	ctx := context.Background()

	{
		shortenerService := NewService(rdb, ctx)
		shortenerGroup.GET("/:shortURL", shortenerService.RedirectToOriginal)
		shortenerGroup.POST("/shorten", shortenerService.ShortenURL)
	}
}
