package models

import "fmt"

func (store SqlStore) CreateNewNote(newNoteRequest *NewNoteRequest) (*Note, error) {
	userId, err := store.GetUserIdByEmail(newNoteRequest.Email)
	if err != nil {
		return &Note{}, err
	}
	// set AuthorId
	newNoteRequest.NewNote.AuthorID = userId
	fmt.Println("author id is: ", userId)
	fmt.Println("newNote: ", newNoteRequest.NewNote)

	err = store.db.Model(&Note{}).Create(&newNoteRequest.NewNote).Error
	if err != nil {
		return &Note{}, err
	}

	return newNoteRequest.NewNote, nil
}
