package internal_test

import (
	"context"
	"fmt"
	ratingsDBStorage "github.com/go-park-mail-ru/2021_1_kekEnd/internal/ratings/repository/dbstorage"
	ratingsUseCase "github.com/go-park-mail-ru/2021_1_kekEnd/internal/ratings/usecase"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestRating(t *testing.T) {
	os.Setenv("DB_CONNECT", "postgres://mdb:mdb@localhost:5432/mdb")

	connStr, connected := os.LookupEnv("DB_CONNECT")
	if !connected {
		fmt.Println(os.Getwd())
		log.Fatal("Failed to read DB connection data")
	}
	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	assert.NoError(t, err)

	t.Run("CreateRating and GetRating", func(t *testing.T) {
		_, err = dbpool.Exec(context.Background(), "TRUNCATE TABLE mdb.movie_rating")
		assert.NoError(t, err)

		ratingsRepo := ratingsDBStorage.NewRatingsRepository(dbpool)
		ratingsUC := ratingsUseCase.NewRatingsUseCase(ratingsRepo)

		err = ratingsUC.CreateRating("user1", "2", 4)
		assert.NoError(t, err)

		rating, err := ratingsUC.GetRating("user1", "2")
		assert.NoError(t, err)

		assert.Equal(t, rating.UserID, "user1")
		assert.Equal(t, rating.MovieID, "2")
		assert.Equal(t, rating.Score, 4)
	})

	t.Run("CreateRating and GetRating", func(t *testing.T) {
		_, err = dbpool.Exec(context.Background(), "TRUNCATE TABLE mdb.movie_rating")
		assert.NoError(t, err)

		ratingsRepo := ratingsDBStorage.NewRatingsRepository(dbpool)
		ratingsUC := ratingsUseCase.NewRatingsUseCase(ratingsRepo)

		err = ratingsUC.CreateRating("user1", "2", 4)
		assert.NoError(t, err)

		ratingsRepo2 := ratingsDBStorage.NewRatingsRepository(dbpool)
		ratingsUC2 := ratingsUseCase.NewRatingsUseCase(ratingsRepo2)

		err = ratingsUC2.CreateRating("user1", "2", 4)
		assert.Error(t, err)

	})

	t.Run("UpdateRating", func(t *testing.T) {
		_, err = dbpool.Exec(context.Background(), "TRUNCATE TABLE mdb.movie_rating")
		assert.NoError(t, err)

		ratingsRepo := ratingsDBStorage.NewRatingsRepository(dbpool)
		ratingsUC := ratingsUseCase.NewRatingsUseCase(ratingsRepo)

		err = ratingsUC.CreateRating("user1", "2", 4)
		assert.NoError(t, err)

		rating, err := ratingsUC.GetRating("user1", "2")
		assert.NoError(t, err)

		assert.Equal(t, rating.UserID, "user1")
		assert.Equal(t, rating.MovieID, "2")
		assert.Equal(t, rating.Score, 4)

		err = ratingsUC.UpdateRating("user1", "2", 8)
		assert.NoError(t, err)

		ratingAfterUpdate, err := ratingsUC.GetRating("user1", "2")
		assert.NoError(t, err)

		assert.Equal(t, ratingAfterUpdate.UserID, "user1")
		assert.Equal(t, ratingAfterUpdate.MovieID, "2")
		assert.Equal(t, ratingAfterUpdate.Score, 8)
	})

	t.Run("DeleteRating", func(t *testing.T) {
		_, err = dbpool.Exec(context.Background(), "TRUNCATE TABLE mdb.movie_rating")
		assert.NoError(t, err)

		ratingsRepo := ratingsDBStorage.NewRatingsRepository(dbpool)
		ratingsUC := ratingsUseCase.NewRatingsUseCase(ratingsRepo)

		err = ratingsUC.CreateRating("user1", "2", 4)
		assert.NoError(t, err)

		rating, err := ratingsUC.GetRating("user1", "2")
		assert.NoError(t, err)

		assert.Equal(t, rating.UserID, "user1")
		assert.Equal(t, rating.MovieID, "2")
		assert.Equal(t, rating.Score, 4)

		err = ratingsUC.DeleteRating("user1", "2")
		assert.NoError(t, err)

		_, err = ratingsUC.GetRating("user1", "2")
		assert.Error(t, err)

	})
}
