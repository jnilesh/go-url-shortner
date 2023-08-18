package modules

import (
	"go-url-shortner/internal/modules/shortener"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, shortenerService *shortener.Service) {
	shortener.RegisterRoutes(r, shortenerService)
}
