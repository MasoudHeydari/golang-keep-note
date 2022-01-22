package controller

import (
	"encoding/json"
	"errors"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (server *Server) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, errors.New("error occurred while parsing request body"))
		return
	}

	newUser.Prepare()
	createdUser, err := server.sqlStore.CreateNewUser(&newUser)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusCreated, createdUser)
}

func (server *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers, err := server.sqlStore.GetAllUsers()
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, allUsers)
}

func (server *Server) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	fetchedUser, err := server.sqlStore.GetUserById(uint32(userId))

	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	respond.JSON(w, http.StatusOK, fetchedUser)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
