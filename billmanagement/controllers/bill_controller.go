package controllers

import (
	"net/http"

	"com.electricity.online.bill/service"
	"github.com/gorilla/mux"
)

var billService *service.BillService

type BillController struct {
}

func (billController BillController) RegisterRoutes(r *mux.Router) {
	r.Handle("/api/admin/bills", http.HandlerFunc(billService.CreateBill)).Methods(http.MethodPost)
	r.Handle("/api/admin/bills/{billNumber}", http.HandlerFunc(billService.GetBill)).Methods(http.MethodGet)
	r.Handle("/api/admin/bills", http.HandlerFunc(billService.GetAllBill)).Methods(http.MethodGet)
}
