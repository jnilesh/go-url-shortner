package shortener

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Service encapsulates the logic for the URL shortener. ok

type Service struct {
	rdb *redis.Client
	ctx context.Context
}

// NewService creates a new shortener service.
func NewService(rdb *redis.Client, ctx context.Context) *Service {
	return &Service{rdb: rdb, ctx: ctx}
}

// RedirectToOriginal resolves a shortened URL to its original and redirects the user.
func (s *Service) RedirectToOriginal(c *gin.Context) {
	shortURL := c.Param("shortURL")

	originalURL, err := s.rdb.Get(s.ctx, shortURL).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve original URL"})
		return
	}

	c.Redirect(http.StatusFound, originalURL)
}

// ShortenURL shortens a given URL and returns the shortened version.
func (s *Service) ShortenURL(c *gin.Context) {
	originalURL := c.PostForm("url")

	shortURL, err := generateShortURL(originalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
		return
	}

	err = s.rdb.Set(s.ctx, shortURL, originalURL, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store short URL in Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// generateShortURL is an internal helper to create a shortened URL from a given original URL.
func generateShortURL(originalURL string) (string, error) {
	hash := fmt.Sprintf("%x", hashCode(originalURL))
	return hash[:8], nil
}

// hashCode is an internal helper to generate a hash code from a string.
func hashCode(s string) uint32 {
	var h uint32
	for _, c := range s {
		h = 31*h + uint32(c)
	}
	return h
}
