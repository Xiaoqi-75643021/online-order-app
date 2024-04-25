package jwt_test

import (
	"fmt"
	"online-ordering-app/internal/config"
	"online-ordering-app/internal/repository"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"

)

var jwtSecret = []byte(config.Cfg.Jwt.Key)

type Claims struct {
	UserID uint `json:"user_id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func TestGenerateToken(t *testing.T) {
	expirationTime := time.Now().Add(24 * time.Hour)
	user, err := repository.FindUserByName("Mike")
	if err != nil {
		t.Fatal(err)
	}
	claims := &Claims{
		UserID: user.UserID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(tokenString)
}

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo1LCJyb2xlIjoiYWRtaW4iLCJleHAiOjE3MTQxMTU0ODR9.HlgOViVJsRRC2IlcumR_-ObGaAzXFHzOEAnt3CZzNNc"
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		t.Fatal(err)
	}
}
