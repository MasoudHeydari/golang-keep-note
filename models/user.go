package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key" json:"user_id"`
	FirstName string    `gorm:"size:255; not null" json:"first_name"`
	LastName  string    `gorm:"size:255; not null" json:"last_name"`
	Email     string    `gorm:"size:255; not null; unique" json:"email"`
	Password  string    `gorm:"size:255; not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CreatedNewUserResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatedUserResponse struct {
	Message string `json:"message"`
}

type UpdateUserRequest struct {
	Email            string `json:"email"`
	PreviousPassword string `json:"old_password"`
	NewPassword      string `json:"new_password"`
}

func (usr *User) Prepare() {
	usr.ID = 0
	usr.FirstName = strings.TrimSpace(usr.FirstName)
	usr.LastName = strings.TrimSpace(usr.LastName)
	usr.Email = strings.TrimSpace(usr.Email)
}

func (usr *User) HashPassword() error {
	byteHashedPass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(byteHashedPass)
	return nil
}

func (usr *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if usr.FirstName == "" {
			return errors.New("required first name")
		}
		if usr.LastName == "" {
			return errors.New("required last name")
		}
		if usr.Password == "" {
			return errors.New("required password")
		}
		if usr.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(usr.Email); err != nil {
			return errors.New("invalid email")
		}

		return nil
	case "login":
		if usr.Password == "" {
			return errors.New("required password")
		}
		if usr.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(usr.Email); err != nil {
			return errors.New("invalid email format")
		}
		return nil

	default:
		// when 'action' is empty
		if usr.FirstName == "" {
			return errors.New("required first name")
		}
		if usr.LastName == "" {
			return errors.New("required last name")
		}
		if usr.Password == "" {
			return errors.New("required password")
		}
		if usr.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(usr.Email); err != nil {
			return errors.New("invalid email format")
		}
		return nil
	}
}

func (uur *UpdateUserRequest) Validate() error {
	if uur.PreviousPassword == "" {
		return errors.New("required previous password")
	}
	if uur.NewPassword == "" {
		return errors.New("required new password")
	}
	if uur.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(uur.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}
