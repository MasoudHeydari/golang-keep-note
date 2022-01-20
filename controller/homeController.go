package controller

import (
	"github.com/MasoudHeydari/golang-keep-note/auth"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, 200, map[string]string{"massage": "welcome to home page"})
	token, err := auth.CreateToken(23)
	if err != nil {
		respond.Error(w, 400, err)
	}

	respond.JSON(w, 200, token)

}
