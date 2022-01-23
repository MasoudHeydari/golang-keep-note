package controller

import (
	"errors"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"net/http"
)

func (server *Server) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := server.tokenMaker.VerifyToken(r)
		if err != nil {
			respond.Error(w, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		next(w, r)
	}
}
