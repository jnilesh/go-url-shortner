package modules

import (
	"go-url-shortner/internal/modules/shortener"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RegisterRoutes(r *gin.Engine, rdb *redis.Client) {
	shortener.RegisterRoutes(r, rdb)

}
