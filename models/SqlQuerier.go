package models

type SqlQuerier interface {
	CreateNewUser(newUser *User) (*User, error)
	IsEmailAlreadyExist(user *User) (bool, error)
	GetAllUsers() (*[]User, error)
	GetUserById(userId uint32) (*User, error)
	GetUserIdByEmail(email string) (uint32, error)

	GetUserByEmail(email string) (*User, error)
	UpdateUserPassword(updateUserRequest *UpdateUserRequest) (message string, err error)

	LoginUser(user *User) error

	CreateNewNote(newNote *NewNoteRequest) (*Note, error)
	GetAllNotes(email string) (*[]Note, error)
	UpdateNote(note *UpdateNoteRequest) (*UpdateNoteResponse, error)
	DeleteANote(noteId int, email string) error
	GetANoteByNoteId(noteId int) (*Note, error)
}
