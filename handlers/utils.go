package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyNumberIsInCorrectFormat(n int) bool {
	str := strconv.Itoa(int(n))
	if len(str[3:]) != 9 {
		return false
	}
	return true
}

func generateSecureToken() string {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func verifyPasswordMatch(str1, str2 string) bool {
	res := strings.Compare(str1, str2)
	if res == 0 {
		return true
	} else {
		return false
	}
}

var (
	key []byte
	t   *jwt.Token
)

func createJWTToken(username string) (string, error) {
	key = []byte("maasai")              // load from .env
	t = jwt.New(jwt.SigningMethodHS256) // create new token

	claims := t.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authorized"] = true
	claims["username"] = username

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

func VerifyToken(tokenString string) (error, interface{}) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return tokenString, nil
	})
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return err, ""
	}

	c := token.Claims.(jwt.MapClaims)
	usr := c["username"]

	// TODO Later : Validate if token is valid
	// if !token.Valid {
	// 	return fmt.Errorf("invalid token")
	// }

	return nil, usr
}
