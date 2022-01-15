package ratings

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/logger"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/ratings"
)

// Handler структура хендлера
type Handler struct {
	useCase ratings.UseCase
	Log     *logger.Logger
}

// NewHandler инициализация нового хендлера
func NewHandler(useCase ratings.UseCase, Log *logger.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		Log:     Log,
	}
}

// ratingData структура оценок
type ratingData struct {
	MovieID string `json:"movie_id"`
	Score   string `json:"score"`
}

// CreateRating создание оценки
func (h *Handler) CreateRating(ctx *gin.Context) {
	ratingData := new(ratingData)
	err := ctx.BindJSON(ratingData)
	if err != nil {
		msg := "Failed to bind request data" + err.Error()
		h.Log.LogWarning(ctx, "ratings", "CreateRating", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	score, err := strconv.Atoi(ratingData.Score)
	if err != nil {
		err := fmt.Errorf("%s", "Failed to cast rating value to number")
		h.Log.LogWarning(ctx, "ratings", "CreateRating", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	err = h.useCase.CreateRating("user1", ratingData.MovieID, score)
	if err != nil {
		h.Log.LogError(ctx, "ratings", "CreateRating", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusCreated)
}

// GetRating получение оценки
func (h *Handler) GetRating(ctx *gin.Context) {
	movieID := ctx.Param("movie_id")

	rating, err := h.useCase.GetRating("user1", movieID)
	if err != nil {
		h.Log.LogWarning(ctx, "ratings", "GetRating", err.Error())
		ctx.AbortWithStatus(http.StatusNotFound) // 404
		return
	}

	ctx.JSON(http.StatusOK, rating)
}

// UpdateRating обновление оценки
func (h *Handler) UpdateRating(ctx *gin.Context) {
	ratingData := new(ratingData)
	err := ctx.BindJSON(ratingData)
	if err != nil {
		msg := "Failed to bind request data" + err.Error()
		h.Log.LogWarning(ctx, "ratings", "UpdateRating", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	score, err := strconv.Atoi(ratingData.Score)
	if err != nil {
		msg := "Failed to cast rating value to number" + err.Error()
		h.Log.LogWarning(ctx, "ratings", "UpdateRating", msg)
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	err = h.useCase.UpdateRating("user1", ratingData.MovieID, score)
	if err != nil {
		h.Log.LogError(ctx, "ratings", "UpdateRating", err)
		ctx.AbortWithStatus(http.StatusInternalServerError) // 500
		return
	}

	ctx.Status(http.StatusOK)
}

// DeleteRating удаление оценки
func (h *Handler) DeleteRating(ctx *gin.Context) {
	movieID := ctx.Param("movie_id")

	err := h.useCase.DeleteRating("user1", movieID)
	if err != nil {
		h.Log.LogWarning(ctx, "ratings", "DeleteRating", err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest) // 400
		return
	}

	ctx.Status(http.StatusOK)
}
