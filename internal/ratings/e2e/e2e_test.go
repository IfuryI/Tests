package e2e_test

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2021_1_kekEnd/internal/server"
	constants "github.com/go-park-mail-ru/2021_1_kekEnd/pkg/const"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestE2ERating(t *testing.T) {
	os.Setenv("DB_CONNECT", "postgres://mdb:mdb@127.0.0.1:5432/mdb")

	connStr, connected := os.LookupEnv("DB_CONNECT")
	if !connected {
		fmt.Println(os.Getwd())
		log.Fatal("Failed to read DB connection data")
	}
	dbpool, err := pgxpool.Connect(context.Background(), connStr)
	assert.NoError(t, err)


	app := server.NewApp()

	go func() {
		err := app.Run(constants.Port)
		assert.NoError(t, err)
	}()

	t.Run("CreateRating", func(t *testing.T) {
		_, err = dbpool.Exec(context.Background(), "TRUNCATE TABLE mdb.movie_rating")
		assert.NoError(t, err)


		// AddRating
		reqStrAddRating := `{"movie_id":"1", "score": "4"}`
		reqAddRating, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%s/api/v1/ratings", constants.Port), strings.NewReader(reqStrAddRating))
		assert.NoError(t, err)

		reqAddRating.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		client := http.Client{}
		response, err := client.Do(reqAddRating)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, response.StatusCode)


		// GetRating
		reqGetRating, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%s/api/v1/ratings/%d", constants.Port, 1), strings.NewReader(""))
		assert.NoError(t, err)

		reqGetRating.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		clientGetRating := http.Client{}
		responseGetRating, err := clientGetRating.Do(reqGetRating)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, responseGetRating.StatusCode)

		byteBodyGetRating, err := ioutil.ReadAll(responseGetRating.Body)
		assert.NoError(t, err)

		assert.Equal(t, `{"username":"user1","movie_id":"1","score":4}`, strings.Trim(string(byteBodyGetRating), "\n"))
		response.Body.Close()


		// UpdateRating
		reqStrUpdateRating := `{"movie_id":"1", "score": "8"}`
		reqUpdateRating, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:%s/api/v1/ratings", constants.Port), strings.NewReader(reqStrUpdateRating))
		assert.NoError(t, err)

		reqUpdateRating.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		clientUpdateRating := http.Client{}
		responseUpdateRating, err := clientUpdateRating.Do(reqUpdateRating)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, responseUpdateRating.StatusCode)


		// GetRatingAfterUpdate
		clientGetRatingAfterUpdate := http.Client{}
		responseGetRatingAfterUpdate, err := clientGetRatingAfterUpdate.Do(reqGetRating)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, responseGetRatingAfterUpdate.StatusCode)

		byteBodyGetRatingAfterUpdate, err := ioutil.ReadAll(responseGetRatingAfterUpdate.Body)
		assert.NoError(t, err)

		assert.Equal(t, `{"username":"user1","movie_id":"1","score":8}`, strings.Trim(string(byteBodyGetRatingAfterUpdate), "\n"))
		response.Body.Close()


		// DeleteRating
		reqDeleteRating, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:%s/api/v1/ratings/%d", constants.Port, 1), strings.NewReader(""))
		assert.NoError(t, err)

		reqDeleteRating.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		clientDeleteRating := http.Client{}
		responseDeleteRating, err := clientDeleteRating.Do(reqDeleteRating)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, responseDeleteRating.StatusCode)


		// GetRatingAfterDelete
		clientGetRatingAfterDelete := http.Client{}
		responseGetRatingAfterDelete, err := clientGetRatingAfterDelete.Do(reqGetRating)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, responseGetRatingAfterDelete.StatusCode)

		byteBodyGetRatingAfterDelete, err := ioutil.ReadAll(responseGetRatingAfterDelete.Body)
		assert.NoError(t, err)

		assert.Equal(t, "", strings.Trim(string(byteBodyGetRatingAfterDelete), "\n"))
		response.Body.Close()
	})
}