package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyNumberIsInCorrectFormat(n int) bool {
	str := strconv.Itoa(int(n))
	return len(str[3:]) == 9
}

// func verifyPasswordMatch(str1, str2 string) bool {
// 	res := strings.Compare(str1, str2)
// 	if res == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

var (
	key []byte
	t   *jwt.Token
)

func createJWTToken(username string, id int) (string, error) {
	key = []byte("maasai")              // load from .env
	t = jwt.New(jwt.SigningMethodHS256) // create new token

	claims := t.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authorized"] = true
	claims["username"] = username
	claims["u_id"] = id

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

func VerifyToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return tokenString, nil
	})
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return "", err
	}

	c := token.Claims.(jwt.MapClaims)
	usr := c["username"]
	fmt.Println(c["u_id"])
	// TODO Later : Validate if token is valid
	// if !token.Valid {
	// 	return fmt.Errorf("invalid token")
	// }

	return usr, nil
}

func enableCors(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json:charset=utf-8")
	ctx.Header("Host", ctx.Request.Host)
	ctx.Header("X-Powered-By", "Golang")
	ctx.Header("Access-Control-Allow-Origin", "https://findmyiphone.vercel.app")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Token, Accept, X-Requested-With")
}
