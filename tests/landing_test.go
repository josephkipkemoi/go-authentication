package tests

import (
	"jk/go-sportsapp/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLandingAPI(t *testing.T) {
	router := server.ConnectServerRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome: PinaclebetAPI Ver: 2.0", w.Body.String())
}
