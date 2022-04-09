package service

import (
	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken() (string, error)
	ValidateToken(token string)
}

type JWTUtil struct {
	key []byte
}

func NewJWTUtil() *JWTUtil {
	return &JWTUtil{}
}

func (j *JWTUtil) GenerateToken() (string, error) {
	return jwt.New(jwt.SigningMethodHS256).SignedString(j.key)
}

func (j *JWTUtil) ValidateToken(token string) {
	//TODO implement me
	panic("implement me")
}
