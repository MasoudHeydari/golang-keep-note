package models

import (
	"errors"
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/utils"
	"golang.org/x/crypto/bcrypt"
)

func (store *SqlStore) LoginUser(user *User) error {
	foundUser, err := store.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	err = utils.CheckPassword(foundUser.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("hashedPassword is not the hash of the given password")
		return errors.New("email or password is incorrect")
	}

	*user = *foundUser
	return nil
}
