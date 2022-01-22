package models

type SqlQuerier interface {
	CreateNewUser(newUser *User) (*User, error)
	IsEmailAlreadyExist(user *User) (bool, error)
}
