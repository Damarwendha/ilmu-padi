package entity

import (
	"time"

	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

// Initialize Struct Payment
type Transaction struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	User_ID     int    `gorm:"index" json:"user_id"`
	Course_ID   int    `gorm:"index" json:"course_id"`
	Amount      int    `json:"amount"`
	Status      string `json:"status"`
	Code        string `json:"code"`
	Payment_URL string `json:"payment_url"`
	User        User
	Course      Course
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
