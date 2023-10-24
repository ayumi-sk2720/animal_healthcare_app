package model

import "time"

// User struct -  basic user model
type Task struct {
	Id    string    `gorm:"primary_key" json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}
