package service

import (
	"encoding/json"
	"net/http"

	// config "com.electricity.online/db"

	pb "com.electricity.online/pb"
	"com.electricity.online/repository"
	"github.com/gorilla/mux"
	"github.com/micro/micro/v3/service/logger"
)

var eventRepository *repository.EventRepository

type EventService struct{}

func (es *EventService) AddEvent(response http.ResponseWriter, request *http.Request) {
	var event *pb.Event
	_ = json.NewDecoder(request.Body).Decode(&event)
	err := eventRepository.CreateEvent(event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Inserted data with Id:" + event.Id)
	json.NewEncoder(response).Encode(event)
}

func (es *EventService) ReadEventById(response http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	event, err := eventRepository.GetEventById(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Fetched Event With Id:" + id)
	json.NewEncoder(response).Encode(event)
}

// GetEvents godoc
// @Summary Get events
// @Description Get a list of all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} pb.Event
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /events [get]
func (es *EventService) GetEvents(response http.ResponseWriter, r *http.Request) {
	events, err := eventRepository.GetEvents()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	logger.Infof("Fetched all Events")
	json.NewEncoder(response).Encode(events)
}

func (es *EventService) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	var event *pb.Event
	_ = json.NewDecoder(request.Body).Decode(&event)
	err := eventRepository.UpdateEvent(event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Updated Event with Id:" + event.Id)
	json.NewEncoder(response).Encode("Updated Event with Id:" + event.Id)
}

func (es *EventService) DeleteEvent(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	err := eventRepository.DeleteEvent(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	logger.Infof("Deleted Event with Id:" + id)
	json.NewEncoder(response).Encode("Deleted Event with Id:" + id)
}
