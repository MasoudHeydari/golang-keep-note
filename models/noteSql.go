package models

import "github.com/jinzhu/gorm"

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
	return nil
}
