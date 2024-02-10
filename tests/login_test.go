package tests

import (
	"bytes"
	"encoding/json"
	"jk/go-sportsapp/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Login struct {
	PhoneNumber int64
	Password    string
}

func TestRegisteredUserCanLogIn(t *testing.T) {
	router := server.ConnectServerRouter()

	u := Login{
		PhoneNumber: 254700545727,
		Password:    "1234",
	}
	w := httptest.NewRecorder()

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(u)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/api/login", &buf)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
