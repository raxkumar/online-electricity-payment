package iam

import (
	"net/http"

	"github.com/gorilla/mux"
)

var iamService *IAMService

type IAMController struct {
}

func (iam IAMController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/iam/users", http.HandlerFunc(iamService.CreateUser)).Methods(http.MethodPost)
	r.Handle("/api/iam/users", http.HandlerFunc(iamService.ListUsers)).Methods(http.MethodGet)
	r.Handle("/api/iam/users/{id}", http.HandlerFunc(iamService.GetUserByID)).Methods(http.MethodGet)
	r.Handle("/api/iam/users/{id}", http.HandlerFunc(iamService.UpdateUserDetails)).Methods(http.MethodPut)
	r.Handle("/api/iam/users/{id}", http.HandlerFunc(iamService.DisableUser)).Methods(http.MethodDelete)
}
