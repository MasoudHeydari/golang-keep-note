package models

import (
	"errors"
	"fmt"
	"github.com/MasoudHeydari/golang-keep-note/utils"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func (store *SqlStore) CreateNewUser(newUser *User) (*User, error) {
	isEmailExist, err := store.IsEmailAlreadyExist(newUser)
	fmt.Println("is email already exist: ", isEmailExist)
	if err != nil {
		return &User{}, err
	}

	if isEmailExist {
		log.Println("user already exist, try another email")
		return &User{}, errors.New("this email already registered, try another one")
	}
	err = newUser.Validate("")
	if err != nil {
		return &User{}, err
	}
	err = newUser.HashPassword()
	if err != nil {
		return &User{}, err
	}
	fmt.Println("try to create new user")
	err = store.db.Create(newUser).Error
	if err != nil {
		return &User{}, err
	}
	fmt.Println("after create: ", newUser.ID)
	return newUser, nil
}

func (store *SqlStore) IsEmailAlreadyExist(user *User) (bool, error) {
	isEmailExist := false
	r := store.db.Model(User{}).Where("email = ?", user.Email).Take(&User{})

	if r.Error != nil && !gorm.IsRecordNotFoundError(r.Error) {
		// reach here when we have error and user want to register for first time
		return false, r.Error
	}

	isEmailExist = r.RowsAffected > 0
	return isEmailExist, nil
}

func (store *SqlStore) GetAllUsers() (*[]User, error) {
	users := []User{}
	err := store.db.Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

func (store *SqlStore) GetUserById(userId uint32) (*User, error) {
	user := User{}
	err := store.db.Model(&User{}).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user not found")
	}

	return &user, nil

}

func (store *SqlStore) GetUserByEmail(email string) (*User, error) {
	fetchedUser := User{}
	r := store.db.Model(User{}).Where("email = ?", email).Take(&fetchedUser)
	//fmt.Println("row affected: ", r.RowsAffected)
	if r.Error != nil {
		return &User{}, r.Error
	}
	return &fetchedUser, nil
}

func (store *SqlStore) GetUserIdByEmail(email string) (uint32, error) {
	fetchedUser := User{}
	r := store.db.Model(User{}).Where("email = ?", email).Take(&fetchedUser)
	if r.Error != nil {
		return fetchedUser.ID, r.Error
	}
	return fetchedUser.ID, nil
}

func (store *SqlStore) UpdateUserPassword(updateUserRequest *UpdateUserRequest) (message string, err error) {
	if err := updateUserRequest.Validate(); err != nil {
		return err.Error(), err
	}

	// hash old password and check it with original password stored in database
	isPasswordsMatch, err := store.isPasswordMatchWithOriginalPassword(updateUserRequest.PreviousPassword, updateUserRequest.Email)
	if err != nil {
		return "failed to update password", err
	}

	if !isPasswordsMatch {
		return "email or password is incorrect", errors.New("email or password is incorrect")
	}

	hashedNewPassword, err := utils.HashPassword(updateUserRequest.NewPassword)
	if err != nil {
		return "failed to update password", err
	}

	dbResult := store.db.Model(&User{}).Where("email = ?", updateUserRequest.Email).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   hashedNewPassword,
			"updated_at": time.Now(),
		})

	if dbResult.Error != nil {
		return "failed to update password", dbResult.Error
	}

	return "user password successfully updated", nil
}

func (store *SqlStore) isPasswordMatchWithOriginalPassword(oldPassword, email string) (bool, error) {
	user, err := store.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Email != email {
		// the email stored in database and the email that user entered, doesn't match
		return false, errors.New("email or password is incorrect")
	}

	err = utils.CheckPassword(user.Password, oldPassword)
	if err != nil {
		return false, err
	}

	// the old password matches with password stored in database
	return true, nil
}
