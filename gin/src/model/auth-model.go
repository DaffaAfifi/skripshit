package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Nama  string `json:"nama"`
	Role  int    `json:"role"`
	jwt.StandardClaims
}
