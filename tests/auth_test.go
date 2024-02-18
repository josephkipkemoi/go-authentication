package tests

import (
	"bytes"
	"encoding/json"
	"jk/go-sportsapp/database"
	"jk/go-sportsapp/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCannotRegisterWithInvalidPassword(t *testing.T) {
	router := server.ConnectServerRouter()

	w := httptest.NewRecorder()

	u := database.User{
		Email:    "700545727",
		Password: "1234",
		Passcode: 12345,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(u)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/api/register", &buf)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCanRegisterUser(t *testing.T) {
	router := server.ConnectServerRouter()

	w := httptest.NewRecorder()

	u := database.User{
		Email:    "700545727",
		Password: "1234",
		Passcode: 1234,
	}

	var buf bytes.Buffer

	err := json.NewEncoder(&buf).Encode(u)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/api/register", &buf)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
