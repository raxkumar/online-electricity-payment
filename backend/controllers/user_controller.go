package controllers

import (
	"net/http"

	"com.electricity.online/service"
	"github.com/gorilla/mux"
)

var userService *service.UserService

type UserController struct {
}

func (userController UserController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/users", http.HandlerFunc(userService.AddUser)).Methods(http.MethodPost)
	r.Handle("/api/users", http.HandlerFunc(userService.GetAllUsers)).Methods(http.MethodGet)
	r.Handle("/api/users/{id}", http.HandlerFunc(userService.GetUserByID)).Methods(http.MethodGet)
	r.Handle("/api/users/{id}", http.HandlerFunc(userService.UpdateUser)).Methods(http.MethodPut)
	r.Handle("/api/users/zones/{zoneId}", http.HandlerFunc(userService.GetUsersByZone)).Methods(http.MethodGet)
}
