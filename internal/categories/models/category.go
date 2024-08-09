package models

type Category struct {
	ID              int64  `json:"id" gorm:"primary_key"`
	Name            string `json:"name" gorm:"unique,not null"`
	IsSystemDefault bool   `json:"is_system_default"`
	CreatedBy       uint   `json:"created_by"`
}
