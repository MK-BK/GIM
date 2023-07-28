package models

import "github.com/dgrijalva/jwt-go"

type LoginClaims struct {
	UserName string
	UserID   string
	jwt.StandardClaims
}
