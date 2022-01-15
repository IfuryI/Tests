package ratings

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/logger"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/middleware"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/ratings"
)

// RegisterHTTPEndpoints Зарегестрировать хендлеры
func RegisterHTTPEndpoints(router *gin.RouterGroup, ratingsUC ratings.UseCase, authMiddleware middleware.Auth,
	Log *logger.Logger) {
	handler := NewHandler(ratingsUC, Log)

	router.POST("/ratings", authMiddleware.CheckAuth(false), handler.CreateRating)
	router.GET("/ratings/:movie_id", authMiddleware.CheckAuth(false), handler.GetRating)
	router.PUT("/ratings", authMiddleware.CheckAuth(false), handler.UpdateRating)
	router.DELETE("/ratings/:movie_id", authMiddleware.CheckAuth(false), handler.DeleteRating)
}
