package controllers

import (
	"net/http"

	"com.electricity.online/service"
	"github.com/gorilla/mux"
)

type TestController struct {
}

// RegisterRoutes registers routes for the TestController
func (controller TestController) RegisterRoutes(router *mux.Router) {
	router.Handle("/test", http.HandlerFunc(service.HandleTestRoute)).Methods("GET")
	router.Handle("/test", http.HandlerFunc(service.TestMethod)).Methods("POST")
}
