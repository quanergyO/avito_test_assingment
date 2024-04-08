package types

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int  `json:"user_id"`
	Role   Role `json:"role"`
}
