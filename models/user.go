package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key; auto_increment; not null" json:"user_id"`
	FirstName string    `gorm:"size:255; not null" json:"first_name"`
	LastName  string    `gorm:"size:255; not null" json:"last_name"`
	Email     string    `gorm:"size:255; not null; unique" json:"email"`
	Password  string    `gorm:"size:255; not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
