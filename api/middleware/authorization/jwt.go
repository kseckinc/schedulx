package authorization

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(signKey string) *JwtSign {
	if len(signKey) <= 0 {
		signKey = "bridgx"
	}
	return &JwtSign{
		[]byte(signKey),
	}
}

type JwtSign struct {
	SigningKey []byte
}

func (j *JwtSign) CreateToken(claims CustomClaims) (string, error) {
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenPartA.SignedString(j.SigningKey)
}

func (j *JwtSign) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if token == nil || err != nil {
		return nil, errors.New("token invalid")
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token invalid")
	}
}

func (j *JwtSign) RefreshToken(tokenString string, extraAddSeconds int64) (string, error) {
	if customClaims, err := j.ParseToken(tokenString); err == nil {
		customClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.CreateToken(*customClaims)
	} else {
		return "", err
	}
}
