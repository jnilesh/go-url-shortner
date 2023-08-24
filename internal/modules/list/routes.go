package list

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RegisterRoutes(r *gin.Engine, rdb *redis.Client) {
	listGroup := r.Group("/")
	ctx := context.Background()

	{
		listService := NewService(rdb, ctx)
		listGroup.GET("/all", listService.GetAllSavedURLs)
		// shortenerGroup.POST("/shorten", svc.ShortenURL)
	}
}
