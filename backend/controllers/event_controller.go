package controllers

import (
	"net/http"

	auth "com.electricity.online/auth"
	"com.electricity.online/service"
	"github.com/gorilla/mux"
)

var eventService *service.EventService

type EventController struct {
}

func (eventController EventController) RegisterRoutes(r *mux.Router) {
	r.Handle("/events", auth.Protect(http.HandlerFunc(eventService.AddEvent), "ROLE_ADMIN")).Methods(http.MethodPost)

	// r.Handle("/events", auth.Protect(http.HandlerFunc(eventService.GetEvents), "ROLE_ADMIN", "ROLE_USER")).Methods(http.MethodGet)
	r.Handle("/events", (http.HandlerFunc(eventService.GetEvents))).Methods(http.MethodGet)

	r.Handle("/events/{id}", http.HandlerFunc(eventService.ReadEventById)).Methods(http.MethodGet)
	r.Handle("/events", auth.Protect(http.HandlerFunc(eventService.UpdateEvent))).Methods(http.MethodPatch)
	// r.Handle("/events/{id}", auth.Protect(http.HandlerFunc(eventService.DeleteEvent))).Methods(http.MethodDelete)
}
