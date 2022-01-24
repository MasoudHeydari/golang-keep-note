package controller

import (
	"errors"
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"net/http"
	"strings"
)

var (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func (server *Server) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		authorizationFields := strings.Fields(authorizationHeader)
		if len(authorizationFields) < 2 {
			err := errors.New("invalid authorization header format")
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(authorizationFields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		accessToken := authorizationFields[1]
		payload, err := server.tokenMaker.VerifyTokenPayload(accessToken)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, err)
			return
		}

		r.Header.Set("email", payload.Email)
		next(w, r)
	}
}
