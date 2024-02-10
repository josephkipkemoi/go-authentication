package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jk/go-sportsapp/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Balance struct {
	amount int
}

func TestCanGetUserBalance(t *testing.T) {
	router := server.ConnectServerRouter()
	userId := 1

	b := Balance{
		amount: 10,
	}
	w := httptest.NewRecorder()

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(b)
	if err != nil {
		log.Fatal(err)
	}
	requestUrl := fmt.Sprintf("/api/users/%d/balance", userId)
	req, _ := http.NewRequest(http.MethodPost, requestUrl, &buf)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
