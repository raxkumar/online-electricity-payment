// service/bill_service.go

package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"com.electricity.online.bill/models"
	"com.electricity.online.bill/repository"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gorilla/mux"
)

var billRepository *repository.BillRepository

type BillService struct{}

func (bs *BillService) CreateBill(response http.ResponseWriter, request *http.Request) {
	var bill *models.Bill
	_ = json.NewDecoder(request.Body).Decode(&bill)

	createdBill, err := billRepository.CreateBill(bill)
	logger.Info(createdBill)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(createdBill)
}

func (bs *BillService) GetBill(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	billNumber := vars["billNumber"]

	foundBill, err := billRepository.GetBillByID(billNumber)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{ "message": "Bill not found" }`))
		return
	}
	json.NewEncoder(response).Encode(foundBill)
}

func (bs *BillService) GetAllBill(response http.ResponseWriter, request *http.Request) {
	paymentStatus := request.URL.Query().Get("paymentStatus")
	if paymentStatus == "" {
		// If no paymentStatus filter, get all bills
		bills, err := billRepository.GetAllBills()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(bills)
	} else {
		// If paymentStatus filter provided, get bills with matching paymentStatus
		paymentStatus = strings.ToUpper(paymentStatus)
		if paymentStatus != "UNPAID" && paymentStatus != "PAID" {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{ "message": "Invalid paymentStatus filter" }`))
			return
		}

		filteredBills, err := billRepository.GetBillsByPaymentStatus(paymentStatus)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(filteredBills)
	}
}
