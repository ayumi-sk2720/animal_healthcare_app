package model

import "time"

// User struct -  basic user model
type Task struct {
	Id    int       `gorm:"primary_key" json:"id"`
	Title string    `json:"name"`
	Date  time.Time `json:"date"`
}
