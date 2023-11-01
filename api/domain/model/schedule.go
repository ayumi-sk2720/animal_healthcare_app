package model

import "time"

type Schedule struct {
	Id        int       `gorm:"primary_key" json:"id"`
	PetId     string    `json:"pet_id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	Location  string    `json:"location" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Date struct {
	From time.Time
	To   time.Time
}

type Date2 struct {
	Year  int
	Month time.Month
	Day   int
}
