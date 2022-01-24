package auth

import (
	"net/http"
	"time"
)

// TokenMaker is an interface for managing tokens
type TokenMaker interface {
	VerifyToken(r *http.Request) error
	CreateToken(userIid uint32) (string, error)
	ExtractToken(r *http.Request) string
	ExtractTokenID(r *http.Request) (uint32, error)

	// payload
	CreateTokenPayload(email string, duration time.Duration) (string, error)
	VerifyTokenPayload(token string) (*Payload, error)
}
