package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

func (store SqlStore) CreateNewNote(newNoteRequest *NewNoteRequest) (*Note, error) {
	userId, err := store.GetUserIdByEmail(newNoteRequest.Email)
	if err != nil {
		return &Note{}, err
	}
	// set AuthorId
	newNoteRequest.NewNote.AuthorID = userId

	err = store.db.Model(&Note{}).Create(&newNoteRequest.NewNote).Error
	if err != nil {
		return &Note{}, err
	}

	return newNoteRequest.NewNote, nil
}

func (store *SqlStore) GetAllNotes(email string) (*[]Note, error) {
	allNotes := []Note{}
	authorId, err := store.GetUserIdByEmail(email)
	if err != nil {
		return &allNotes, err
	}

	err = store.db.Model(&Note{}).Where("author_id = ?", authorId).Find(&allNotes).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return &allNotes, err
	}
	return &allNotes, nil
}

func (store *SqlStore) DeleteANote(noteId int, email string) error {
	userId, err := store.GetUserIdByEmail(email)
	if err != nil {
		return errors.New("internal server error, failed to delete note")
	}

	// get note by id
	note, err := store.GetANoteByNoteId(noteId)
	if err != nil {
		return errors.New("internal server error, failed to delete note")
	}

	// check if authenticated user is owner of note
	if note.AuthorID != userId {
		return errors.New("cannot delete this note! your are not the owner of this note")
	}

	//fmt.Println("note: ", note.AuthorID, userId)

	result := store.db.Model(&Note{}).Where("id = ? and author_id = ?", noteId, userId).Take(&Note{}).Delete(&Note{})
	if result.Error != nil {
		return errors.New("internal server error, failed to delete note")
	}

	// note deleted successfully
	return nil
}

func (store *SqlStore) GetANoteByNoteId(noteId int) (*Note, error) {
	fetchedNote := Note{}
	result := store.db.Model(&Note{}).Where("id = ?", noteId).Take(&fetchedNote)
	if result.Error != nil {
		return &Note{}, result.Error
	}
	return &fetchedNote, nil
}
