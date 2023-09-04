package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	customers := jwt.MapClaims{
		"name":   "subasri",
		"rollno": 304,
		"iat":    time.Now().Unix(),                // Issued at timestamp
		"exp":    time.Now().Add(time.Hour).Unix(), // Expiration timestamp
	}
	secretKey := []byte("Secret key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customers)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JWT Token:")
    fmt.Println(tokenString)
}
