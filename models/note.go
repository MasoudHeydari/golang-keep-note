package models

import "time"

type Note struct {
	ID        uint64    `gorm:"primary_key; auto_increment; not null" json:"note_id"`
	Title     string    `gorm:"size:255;not null;" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	AuthorID  uint32    `sql:"type:int REFERENCES user(user_id)" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
