package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const minSecretKey = 32

type JWTMaker struct {
	secretKey string
}

func NewJWtMaker(secretKey string) (TokenMaker, error) {
	if len(secretKey) < minSecretKey {
		return nil, fmt.Errorf("invalid secret key size: your key size must be at least %d characters", minSecretKey)
	}

	return &JWTMaker{secretKey}, nil
}

func (jwtMaker *JWTMaker) CreateToken(userName string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userName, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(jwtMaker.secretKey))
}

func (jwtMaker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(jwtMaker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrorInvalidToken
	}

	return payload, nil
}
