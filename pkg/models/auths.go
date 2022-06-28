package models

import "github.com/dgrijalva/jwt-go"

type LoginModel struct {
	Phone int32  `json:"phone"`
	Pass  string `json:"password"`
}

type RegisterModel struct {
	Phone int32  `json:"phone"`
	Pass  string `json:"password"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserPhone int32 `json:"userPhone"`
}
