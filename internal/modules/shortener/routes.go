package shortener

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, svc *Service) {
	shortenerGroup := r.Group("/")
	{
		shortenerGroup.GET("/:shortURL", svc.RedirectToOriginal)
		shortenerGroup.POST("/shorten", svc.ShortenURL)
	}
}
