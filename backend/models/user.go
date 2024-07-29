package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents the user model
// User struct with all fields using pointers
type User struct {
	ID         uuid.UUID `gorm:"primary_key" json:"id,omitempty"`
	FirstName  *string   `gorm:"column:first_name" json:"firstName,omitempty"`
	LastName   *string   `gorm:"column:last_name" json:"lastName,omitempty"`
	ZoneID     *string   `gorm:"column:zone_id" json:"zoneID,omitempty"`
	Email      *string   `gorm:"column:email;uniqueIndex" json:"email,omitempty"`
	Phone      *string   `gorm:"column:phone;uniqueIndex" json:"phone,omitempty"`
	Street     *string   `gorm:"column:street" json:"street,omitempty"`
	City       *string   `gorm:"column:city" json:"city,omitempty"`
	PostalCode *string   `gorm:"column:postal_code" json:"postalCode,omitempty"`
	IAMID      *string   `gorm:"column:iam_id;uniqueIndex" json:"IAMID,omitempty"`
	gorm.Model `json:"-"`
}

// BeforeSave hook to validate uniqueness only for non-empty values

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// Assign UUID if ID is not set
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	// if tx.Statement.Changed("Email") {
	// 	var count int64
	// 	tx.Model(&User{}).Where("email = ?", u.Email).Not("id = ?", u.ID).Count(&count)
	// 	if count > 0 {
	// 		return gorm.ErrDuplicatedKey
	// 	}
	// }

	// if tx.Statement.Changed("Phone") {
	// 	var count int64
	// 	tx.Model(&User{}).Where("phone = ?", u.Phone).Not("id = ?", u.ID).Count(&count)
	// 	if count > 0 {
	// 		return gorm.ErrDuplicatedKey
	// 	}
	// }

	// if tx.Statement.Changed("IAMID") {
	// 	var count int64
	// 	tx.Model(&User{}).Where("iam_id = ?", u.IAMID).Not("id = ?", u.ID).Count(&count)
	// 	if count > 0 {
	// 		return gorm.ErrDuplicatedKey
	// 	}
	// }

	return nil
}
