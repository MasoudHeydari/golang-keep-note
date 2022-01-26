package controller

import (
	"github.com/gorilla/mux"
)

func (server *Server) initializeRouters() {
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
	server.Router.HandleFunc("/users", server.CreateNewUser).Methods("POST")
	server.Router.HandleFunc("/users", server.GetAllUsers).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.GetUserById).Methods("GET")
	server.Router.HandleFunc("/users", server.AuthMiddleware(server.UpdateUserPassword)).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", server.AuthMiddleware(server.DeleteUser)).Methods("DELETE")
}

func initializeLoginRouter(server *Server) {
	server.Router.HandleFunc("/login", server.Login).Methods("POST")
}

func initializeNoteRouter(server *Server) {
	server.Router.HandleFunc("/posts", server.AuthMiddleware(server.CreateNewNote)).Methods("POST")
	server.Router.HandleFunc("/posts", server.AuthMiddleware(server.GetAllNotes)).Methods("GET")
}
