package models

import (
	"errors"
	"log"
)

func (store SqlStore) CreateNewUser(newUser *User) (*User, error) {
	isEmailExist, err := store.IsEmailAlreadyExist(newUser)
	if isEmailExist && err != nil {
		log.Println("user already exist, try another email")
		errMsg := err.Error() + " | user already exist, try another email"
		return &User{}, errors.New(errMsg)
	}

	err = newUser.Validate("")
	if err != nil {
		return &User{}, err
	}

	err = store.db.Create(newUser).Error
	if err != nil {
		return &User{}, err
	}

	return newUser, nil
}

func (store *SqlStore) IsEmailAlreadyExist(user *User) (bool, error) {
	var isEmailExist bool
	err := store.db.Model(user).
		Select("COUNT(*) > 0").
		Where("email = ?", user.Email).
		Find(&isEmailExist).
		Error
	return isEmailExist, err
}
