package jwt

import (
	"bluebell/conf"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func GenToken(userID int64, username string) (string, error) {
	//logger.Log.Info(time.Now().Add(time.Hour * time.Duration(conf.Conf.Auth.Exp)))
	c := MyClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Jannan",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(conf.Conf.Auth.Exp))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(conf.Conf.Auth.Secret))
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
