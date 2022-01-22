package models

import (
	"errors"
	"fmt"
	"log"
)

func (store *SqlStore) CreateNewUser(newUser *User) (*User, error) {
	isEmailExist, err := store.IsEmailAlreadyExist(newUser)
	fmt.Println("is email already exist: ", isEmailExist)
	if isEmailExist && err != nil {
		log.Println("user already exist, try another email")
		errMsg := err.Error() + " | user already exist, try another email"
		return &User{}, errors.New(errMsg)
	}
	fmt.Println("before valid: ", newUser.ID)
	err = newUser.Validate("")
	if err != nil {
		return &User{}, err
	}
	fmt.Println("after valid: ", newUser.ID)

	err = store.db.Create(newUser).Error
	if err != nil {
		return &User{}, err
	}
	fmt.Println("after create: ", newUser.ID)
	return newUser, nil
}

func (store *SqlStore) IsEmailAlreadyExist(user *User) (bool, error) {
	// TODO this function has bug, fix it... MasoudHeydari: 2022 jan 22 - 22:31
	var isEmailExist bool
	fmt.Println("email: ", user.Email)
	r := store.db.Model(&User{}).
		Select("COUNT(*) > 0").
		Where("email = ?", user.Email).
		Scan(&isEmailExist)

	r = store.db.Raw("select email from users where email = ?", user.Email).Scan(&isEmailExist)
	fmt.Println("rows affected: ", r.RowsAffected)

	fmt.Println("isEmailExist: ", isEmailExist)
	return isEmailExist, r.Error
}

func (store *SqlStore) GetAllUsers() (*[]User, error) {
	users := []User{}
	err := store.db.Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, nil
}
