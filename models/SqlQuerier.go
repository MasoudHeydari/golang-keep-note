package models

type SqlQuerier interface {
	CreateNewUser(newUser *User) (*User, error)
	IsEmailAlreadyExist(user *User) (bool, error)
	GetAllUsers() (*[]User, error)
	GetUserById(userId uint32) (*User, error)

	GetUserByEmail(email string) (*User, error)
	UpdateUserPassword(updateUserRequest *UpdateUserRequest) (message string, err error)

	LoginUser(user *User) error
}
