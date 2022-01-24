package controller

import (
	"encoding/json"
	"errors"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	createdUserResponse := buildCreatedNewUserResponse(createdUser)
	respond.JSON(w, http.StatusCreated, createdUserResponse)
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

func (server *Server) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUserRequest := models.UpdateUserRequest{}
	err = json.Unmarshal(body, &updatedUserRequest)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUserRequest.Email = r.Header.Get("email")

	msg, err := server.sqlStore.UpdateUserPassword(&updatedUserRequest)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, buildUpdatedUserResponse(msg))
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, "will be implemented soon...")
}

func buildCreatedNewUserResponse(user *models.User) *models.CreatedNewUserResponse {
	return &models.CreatedNewUserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func buildUpdatedUserResponse(message string) *models.UpdatedUserResponse {
	return &models.UpdatedUserResponse{
		Message: message,
	}
}
