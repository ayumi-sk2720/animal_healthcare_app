package main_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/myantyuWorld/animal_healthcate/handler"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
}

func TestAPI(t *testing.T) {
	h := &handler.PetHandler{
		//mock入れる
	}

	e := echo.New()
	handler.InitRouting(e, *h)

	t.Run("get", func(t *testing.T) {
		var req *http.Request
		rec := httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodGet, "pet/1", nil)
		e.ServeHTTP(rec, req)

		assert.Equal(t, req.Response.StatusCode, http.StatusOK)
	})
}
