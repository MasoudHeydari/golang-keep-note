package controller

import (
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/database"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB       *gorm.DB
	sqlStore models.SqlQuerier
	Router   *mux.Router
}

var server = Server{}

func (server *Server) InitializeDB() {
	dbConnection, err := database.ConnectToDB()
	if err != nil {
		log.Fatal("failed to connect to database, error: ", err)
	}
	server.sqlStore = models.NewSqlStore(dbConnection)
	server.DB = dbConnection
}

func (server *Server) Run(address string) {
	fmt.Println("Listening to port: 8000")
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func Run() {
	server.InitializeDB()
	server.InitializeRouters()
	server.Run("localhost:8000")
}
