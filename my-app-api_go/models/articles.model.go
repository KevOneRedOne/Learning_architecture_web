package models

type Article struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Description string `json:"description"`
	Date string `json:"date"`
}