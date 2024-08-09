package models

type Expense struct {
	ID          int64   `json:"id" gorm:"primary_key"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Description string  `json:"description"`
	CategoryID  int64   `json:"category_id" gorm:"not null"`
	CreatedBy   int64   `json:"created_by"`
}
