package controller

import (
	"encoding/json"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"github.com/MasoudHeydari/golang-keep-note/utils"
	"golang.org/x/crypto/bcrypt"
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

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	loginResponse := newLoginResponse(token, &user)
	respond.JSON(w, http.StatusOK, loginResponse)

}

func (server *Server) SignIn(email string, password string) (string, error) {
	user, err := server.sqlStore.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = utils.CheckPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", nil
	}

	return server.tokenMaker.CreateToken(user.ID)
}

func newLoginResponse(accessToken string, user *models.User) loginUserResponse {
	return loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
}

func newUserResponse(user *models.User) userResponse {
	return userResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
