package controller

import (
	"encoding/json"
	"errors"
	"github.com/MasoudHeydari/golang-keep-note/models"
	"github.com/MasoudHeydari/golang-keep-note/respond"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (server *Server) CreateNewNote(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	newNote := models.Note{}
	err = json.Unmarshal(body, &newNote)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	newNote.Prepare()
	err = newNote.Validate()
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userEmail := r.Header.Get("email")
	newNoteRequest := newNote.CreateNewNoteRequest(userEmail)

	// save new note into db
	savedNote, err := server.sqlStore.CreateNewNote(newNoteRequest)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, savedNote)
}

func (server *Server) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	allNotes, err := server.sqlStore.GetAllNotes(email)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.JSON(w, http.StatusOK, allNotes)
}

func (server *Server) UpdateANote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		log.Println(err)
		respond.Error(w, http.StatusBadRequest, errors.New("invalid url syntax, note id not valid"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	noteUpdateRequest := models.UpdateNoteRequest{}
	err = json.Unmarshal(body, &noteUpdateRequest)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	email := r.Header.Get("email")
	err = noteUpdateRequest.Prepare(noteId, email)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// update note
	updateNoteResponse, err := server.sqlStore.UpdateNote(&noteUpdateRequest)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	respond.JSON(w, http.StatusOK, updateNoteResponse)
}

func (server *Server) DeleteANote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		respond.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	email := r.Header.Get("email")
	err = server.sqlStore.DeleteANote(int(noteId), email)
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}

	deleteNoteResponse := map[string]string{
		"message": "note deleted successfully",
	}

	respond.JSON(w, http.StatusOK, deleteNoteResponse)
}
