package controllers

import (
	"encoding/json"
	"net/http"

	auth "com.electricity.online/auth"
	eureka "com.electricity.online/eurekaregistry"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gorilla/mux"
)

type CommunicationController struct {
}

func (communicationController CommunicationController) RegisterRoutes(r *mux.Router) {
	r.Handle("/rest/services/backend", auth.Protect(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		logger.Infof("response sent")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}))).Methods(http.MethodGet)

	r.HandleFunc("/api/services/billmanagement", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "billmanagement") }).Methods(http.MethodGet)
	r.HandleFunc("/api/services/payment", func(w http.ResponseWriter, r *http.Request) { eureka.Client(w, r, "payment") }).Methods(http.MethodGet)
}
