package untils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("最大的敌人就是增删查改orz")

func GenToken(uid string) (string, error) {
	claims := MyClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "TheGodOfThisProject",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
