package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key; auto_increment; not null" json:"user_id"`
	FirstName string    `gorm:"size:255; not null" json:"first_name"`
	LastName  string    `gorm:"size:255; not null" json:"last_name"`
	Email     string    `gorm:"size:255; not null; unique" json:"email"`
	Password  string    `gorm:"size:255; not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (usr *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
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
