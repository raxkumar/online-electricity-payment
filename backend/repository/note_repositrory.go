package repository

import (
	config "com.electricity.online/db"
	pb "com.electricity.online/pb"
)

var notestableName = "notes"

type NoteRepository struct{}

type Note struct {
	Id          uint64 `gorm:"primaryKey" json:"id,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Description string `json:"description,omitempty"`
}

func (nr *NoteRepository) GetNotes() ([]*pb.NotesResponse, error) {
	var notes []*pb.NotesResponse
	err := config.DatabaseClient.Table(notestableName).Find(&notes).Error
	return notes, err
}

func (nr *NoteRepository) CreateNote(note *pb.NotesRequest) error {
	noteData := Note{
		Subject:     note.Subject,
		Description: note.Description,
	}
	return config.DatabaseClient.Table(notestableName).Create(&noteData).Error
}

func (nr *NoteRepository) GetNoteById(id string) (*pb.NotesResponse, error) {
	var note *pb.NotesResponse
	err := config.DatabaseClient.Table(notestableName).First(&note, id).Error
	return note, err
}

func (nr *NoteRepository) UpdateNote(note *pb.NotesResponse) error {
	var existingNote *pb.NotesResponse
	err := config.DatabaseClient.Table(notestableName).First(&existingNote, note.Id).Error
	if err != nil {
		return err
	}
	return config.DatabaseClient.Table(notestableName).Model(&existingNote).Updates(note).Error
}

func (nr *NoteRepository) DeleteNote(id string) error {
	var note *pb.NotesResponse
	result := config.DatabaseClient.Table(notestableName).Where("id = ?", id).Delete(&note)
	return result.Error
}
