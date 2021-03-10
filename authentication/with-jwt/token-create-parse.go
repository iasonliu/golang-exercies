package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("AllYourBase")
	// Create the Claims
	claims := MyCustomClaims{
		"user01",
		"user01@email.com",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
			Issuer:    "testIssuser",
		},
	}
	ss := createToken(mySigningKey, claims)
	fmt.Printf("JWT TOKEN: %v\n", ss)
	validClaims := parseToken(mySigningKey, ss)

	fmt.Printf("Parse JWT TOKEN: %#v\n", validClaims)
}

func createToken(key []byte, claims MyCustomClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		log.Fatalln(err)
	}
	return ss
}

func parseToken(key []byte, tokenString string) *MyCustomClaims {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if token.Valid && err == nil {
		claims := token.Claims.(*MyCustomClaims)
		return claims
	}
	return &MyCustomClaims{}
}
