package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"time"
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

func (store *SqlStore) UpdateNote(note *UpdateNoteRequest) (*UpdateNoteResponse, error) {
	userId, err := store.GetUserIdByEmail(note.Email)
	if err != nil {
		return &UpdateNoteResponse{}, err
	}

	noteToUpdate, err := store.GetANoteByNoteId(int(note.ID))
	if gorm.IsRecordNotFoundError(err) {
		return &UpdateNoteResponse{}, errors.New("failed to update! there is no such note")

	} else if err != nil {
		return &UpdateNoteResponse{}, err
	}

	// check if authenticated user is owner of this note or not
	if userId != noteToUpdate.AuthorID {
		return &UpdateNoteResponse{}, errors.New("cannot update this note! your are not the owner of this note")
	}

	//update note
	noteToUpdate.Title = note.Title
	noteToUpdate.Content = note.Content
	noteToUpdate.UpdatedAt = time.Now()

	result := store.db.Model(&Note{}).Where("id = ?", noteToUpdate.ID).Update(noteToUpdate)
	if result.Error != nil {
		return &UpdateNoteResponse{}, errors.New("failed to update note, internal error")
	}

	log.Println("note updated successfully")
	return noteToUpdate.CreateUpdateNoteResponse(), nil

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
