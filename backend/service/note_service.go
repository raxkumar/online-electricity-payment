package service

import (
	"encoding/json"
	"net/http"

	pb "com.electricity.online/pb"
	"com.electricity.online/repository"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/proto"
)

var notestableName = "notes"
var noteRepository *repository.NoteRepository

type NoteService struct{}

func (ns *NoteService) AddNote(response http.ResponseWriter, request *http.Request) {

	var note *pb.NotesRequest

	_ = json.NewDecoder(request.Body).Decode(&note)
	logger.Info(note)

	err := noteRepository.CreateNote(note)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Inserted data")

	// Encode the response data into binary Protocol Buffers format
	responseData, err := proto.Marshal(note)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "Error encoding response data" }`))
		return
	}
	logger.Info("----", responseData)

	// Set the appropriate headers for binary Protocol Buffers
	response.Header().Set("Content-Type", "application/x-protobuf")

	// Write the binary Protocol Buffers data to the response
	response.Write(responseData)
	// json.NewEncoder(response).Encode(note)
}

func (ns *NoteService) ReadNoteById(response http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	note, err := noteRepository.GetNoteById(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Fetched Note With Id:" + id)
	json.NewEncoder(response).Encode(note)
}

func (ns *NoteService) GetNotes(response http.ResponseWriter, r *http.Request) {
	notes, err := noteRepository.GetNotes()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Fetched all Notes")
	json.NewEncoder(response).Encode(notes)
}

func (ns *NoteService) UpdateNote(response http.ResponseWriter, request *http.Request) {
	var note *pb.NotesResponse
	_ = json.NewDecoder(request.Body).Decode(&note)
	err := noteRepository.UpdateNote(note)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Updated Note with Id:" + note.Id)
	json.NewEncoder(response).Encode("Updated Note with Id:" + note.Id)
}

func (ns *NoteService) DeleteNote(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	err := noteRepository.DeleteNote(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Deleted Note with Id:" + id)
	json.NewEncoder(response).Encode("Deleted Note with Id:" + id)
}
