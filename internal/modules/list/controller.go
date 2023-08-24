package list

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Service encapsulates the logic for the URL shortener.
type Service struct {
	rdb *redis.Client
	ctx context.Context
}

// NewService creates a new shortener service.
func NewService(rdb *redis.Client, ctx context.Context) *Service {
	return &Service{rdb: rdb, ctx: ctx}
}

// GetAllSavedURLs retrieves all saved URLs and returns them as a map of short URLs to original URLs.
func (s *Service) GetAllSavedURLs(c *gin.Context) {
	keys, err := s.rdb.Keys(s.ctx, "*").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve saved URLs"})
		return
	}

	urlMap := make(map[string]string)

	for _, key := range keys {
		originalURL, err := s.rdb.Get(s.ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve original URL"})
			return
		}
		urlMap[key] = originalURL
	}

	c.JSON(http.StatusOK, urlMap)
}
