package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Bill represents the bill model
type Bill struct {
	BillNumber    *string    `gorm:"primary_key;column:bill_number" json:"billNumber,omitempty"`
	ConsumerId    *string    `gorm:"column:consumer_id" json:"consumerId,omitempty"`
	BillDate      *time.Time `gorm:"column:bill_date" json:"billDate,omitempty"`
	DueDate       *time.Time `gorm:"column:due_date" json:"dueDate,omitempty"`
	UnitsConsumed *float64   `gorm:"column:units_consumed" json:"unitsConsumed,omitempty"`
	TotalAmount   *float64   `gorm:"column:total_amount" json:"totalAmount,omitempty"`
	PaymentStatus *string    `gorm:"column:payment_status" json:"paymentStatus,omitempty"`
	UserId        *string    `gorm:"column:user_id" json:"userId,omitempty"`
}

// BeforeSave hook to set default values
func (b *Bill) BeforeSave(tx *gorm.DB) (err error) {
	// Extract month and year from the new bill's bill_date
	newMonth, newYear := b.BillDate.Month(), b.BillDate.Year()
	// Check if there are any existing bills for the same user with the same month and year
	// NOTE: Database-specific Functions: The use of EXTRACT in the query might be database-specific.
	//       If you plan to switch databases, ensure that the date extraction functions are compatible.
	var count int64
	if err := tx.Model(&Bill{}).
		Where("user_id = ? AND EXTRACT(MONTH FROM bill_date) = ? AND EXTRACT(YEAR FROM bill_date) = ?", b.UserId, newMonth, newYear).
		Count(&count).Error; err != nil {
		return err
	}
	// If there are existing bills, return an error
	if count > 0 {
		return errors.New("cannot create more than one bill for the same user in the same month")
	}

	billNumber := generateUniqueBillNumber(b.ConsumerId)
	b.BillNumber = &billNumber
	return nil
}

func generateUniqueBillNumber(consumerID *string) string {
	if consumerID == nil {
		// Handle the case where consumerID is nil
		return ""
	}
	datePrefix := time.Now().Format("20060102")
	randomUUID := uuid.New()
	randomString := fmt.Sprintf("%05d", randomUUID.ID())
	if len(randomString) > 5 {
		randomString = randomString[len(randomString)-5:]
	}
	billNumber := fmt.Sprintf("%s%s%s", datePrefix, *consumerID, randomString)
	return billNumber
}
