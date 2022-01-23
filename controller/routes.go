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
	server.Router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
}

func initializeLoginRouter(server *Server) {

}

func initializeNoteRouter(server *Server) {

}
