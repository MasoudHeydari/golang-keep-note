package models

import (
	"errors"
	"strings"
	"time"
)

type Note struct {
	ID        uint32    `gorm:"primary_key; auto_increment; not null" json:"note_id"`
	Title     string    `gorm:"size:255;not null;" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	AuthorID  uint32    `sql:"type:int REFERENCES user(user_id)" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UpdateNoteRequest struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Email   string
}

type UpdateNoteResponse struct {
	ID        uint32    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewNoteRequest struct {
	Email   string
	NewNote *Note
}

func (note *Note) Prepare() {
	note.ID = 0
	note.Title = strings.TrimSpace(note.Title)
	note.Content = strings.TrimSpace(note.Content)
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
}

func (note *Note) Validate() error {
	if note.Title == "" {
		return errors.New("required title")
	}
	if note.Content == "" {
		return errors.New("required content")
	}

	return nil
}

func (note *Note) CreateNewNoteRequest(email string) *NewNoteRequest {
	return &NewNoteRequest{
		Email:   email,
		NewNote: note,
	}
}

func (note *UpdateNoteRequest) Prepare(noteId int64, email string) error {
	note.ID = uint32(noteId)
	note.Email = email

	if note.Title == "" {
		return errors.New("required title")
	}
	if note.Content == "" {
		return errors.New("required content")
	}
	return nil

}

func (note *Note) CreateUpdateNoteResponse() *UpdateNoteResponse {
	return &UpdateNoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}
