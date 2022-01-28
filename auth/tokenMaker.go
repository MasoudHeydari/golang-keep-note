package auth

import (
	"time"
)

// TokenMaker is an interface for managing tokens
type TokenMaker interface {
	CreateTokenPayload(email string, duration time.Duration) (string, error)
	VerifyTokenPayload(token string) (*Payload, error)
}
