package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type dummyHandler struct{}

func (d *dummyHandler) GetRandomPokemon(c *gin.Context) {
	c.String(200, "dummy")
}

func TestSetupRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := &dummyHandler{}
	SetupRoutes(r, h)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/pokemon/random", nil)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "dummy" {
		t.Fatalf("expected dummy, got %s", w.Body.String())
	}
}
