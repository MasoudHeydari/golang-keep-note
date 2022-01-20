package controller

import (
	"github.com/gorilla/mux"
)

func (server *Server) InitializeRouters() {
	server.Router = mux.NewRouter()
	initializeHomeRouter(server)
	initializeUserRouter(server)
	initializeLoginRouter(server)
	initializeNoteRouter(server)
}

func initializeHomeRouter(server *Server) {
	server.Router.HandleFunc("/", server.Home)
}

func initializeUserRouter(server *Server) {

}

func initializeLoginRouter(server *Server) {

}

func initializeNoteRouter(server *Server) {

}
