package controller

import (
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/auth"
	"github.com/MasoudHeydari/golang-keep-note/config"
	"github.com/MasoudHeydari/golang-keep-note/database"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	sqlStore   models.SqlQuerier
	Router     *mux.Router
	tokenMaker auth.TokenMaker
}

func NewServer() *Server {
	server := Server{}
	server.initializeDB()
	server.initializeRouters()
	server.initializeTokenMaker()
	return &server
}

func (server *Server) initializeDB() {
	dbConnection, err := database.ConnectToDB()
	if err != nil {
		log.Fatal("failed to connect to database, error: ", err)
	}
	server.sqlStore = models.NewSqlStore(dbConnection)
}

func (server *Server) initializeTokenMaker() {
	tokenMaker, err := auth.NewJWtMaker(config.GetTokenSecretKey())
	if err != nil {
		log.Fatal("cannot create token maker, error -> ", err)
	}
	server.tokenMaker = tokenMaker
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port: 8000")
	log.Fatal(http.ListenAndServe(address, server.Router))
}
