package controller

import (
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, 200, map[string]string{"massage": "welcome to home page"})
}
