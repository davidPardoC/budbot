package models

type TransactionType string

const (
	Expense TransactionType = "expense"
	Income  TransactionType = "income"
)

type Transactions struct {
	ID          int64           `json:"id" gorm:"primary_key"`
	Amount      float64         `json:"amount" gorm:"not null"`
	Description string          `json:"description"`
	Type        TransactionType `json:"type"`
	CreatedBy   int64           `json:"created_by"`
}
