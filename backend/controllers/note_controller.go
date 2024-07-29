package controllers

import (
	"net/http"

	auth "com.electricity.online/auth"
	"com.electricity.online/service"
	"github.com/gorilla/mux"
)

var noteService *service.NoteService

type NoteController struct {
}

func (noteController NoteController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/notes", http.HandlerFunc(noteService.AddNote)).Methods(http.MethodPost, http.MethodOptions)
	r.Handle("/api/notes", auth.Protect(http.HandlerFunc(noteService.GetNotes))).Methods(http.MethodGet, http.MethodOptions)
	// r.Handle("/api/notes/{id}", http.HandlerFunc(noteService.ReadNoteById)).Methods(http.MethodGet, http.MethodOptions)
	// r.Handle("/api/notes", auth.Protect(http.HandlerFunc(noteService.UpdateNote))).Methods(http.MethodPatch, http.MethodOptions)
	r.Handle("/api/notes", http.HandlerFunc(noteService.DeleteNote)).Methods(http.MethodDelete, http.MethodOptions)

}
