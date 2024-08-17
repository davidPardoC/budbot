package models

type transactionType string

const (
	Expense transactionType = "expense"
	Income  transactionType = "income"
)

type Transactions struct {
	ID          int64           `json:"id" gorm:"primary_key"`
	Amount      float64         `json:"amount" gorm:"not null"`
	Description string          `json:"description"`
	CategoryID  int64           `json:"category_id"`
	Type        transactionType `json:"type"`
	CreatedBy   int64           `json:"created_by"`
}
