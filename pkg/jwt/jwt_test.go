package jwt_test

import (
	"fmt"
	"online-ordering-app/internal/config"
	"online-ordering-app/internal/repository"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(config.Cfg.Jwt.Key)

type Claims struct {
	UserID uint `json:"userid"`
	jwt.StandardClaims
}

func TestGenerateToken(t *testing.T) {
	expirationTime := time.Now().Add(24 * time.Hour)
	user, err := repository.FindUserByName("mike")
	if err != nil {
		t.Fatal(err)
	}
	claims := &Claims{
		UserID: user.UserID,
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
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjMsImV4cCI6MTcxMjIxNTAxMn0.Tdx6b6XL3Ec8FZUuPwYMH76nvyM3hNzM_hP6QPkzoxg"
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		t.Fatal(err)
	}
}
