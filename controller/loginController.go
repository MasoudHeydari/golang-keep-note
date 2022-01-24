package controller

import (
	"encoding/json"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"io/ioutil"
	"net/http"
	"time"
)

type userResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("login")
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = server.sqlStore.LoginUser(&user)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.tokenMaker.CreateTokenPayload(user.Email, time.Hour)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	loginResponse := buildNewLoginResponse(token, &user)
	respond.JSON(w, http.StatusOK, loginResponse)

}

func buildNewLoginResponse(accessToken string, user *models.User) loginUserResponse {
	return loginUserResponse{
		AccessToken: accessToken,
		User:        buildNewUserResponse(user),
	}
}

func buildNewUserResponse(user *models.User) userResponse {
	return userResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
