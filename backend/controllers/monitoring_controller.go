package controllers

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MonitoringController struct {
}

func (monitoringController MonitoringController) RegisterRoutes(r *mux.Router) {
	r.Handle("/metrics", promhttp.Handler())
}
