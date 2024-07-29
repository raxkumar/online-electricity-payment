// repository/bill_repository.go

package repository

import (
	config "com.electricity.online.bill/db"
	"com.electricity.online.bill/models"
	"github.com/asim/go-micro/v3/logger"
)

type BillRepository struct{}

func (br *BillRepository) CreateBill(bill *models.Bill) (*models.Bill, error) {

	logger.Info(bill)
	if err := config.DatabaseClient.Create(&bill).Error; err != nil {
		return nil, err
	}
	return bill, nil
}

func (br *BillRepository) GetBillByID(billNumber string) (*models.Bill, error) {
	var bill models.Bill
	if err := config.DatabaseClient.Where("bill_number = ?", billNumber).First(&bill).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

func (br *BillRepository) GetAllBills() ([]*models.Bill, error) {
	var bills []*models.Bill
	if err := config.DatabaseClient.Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (br *BillRepository) GetBillsByPaymentStatus(paymentStatus string) ([]*models.Bill, error) {
	var bills []*models.Bill
	if err := config.DatabaseClient.Where("payment_status = ?", paymentStatus).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}
