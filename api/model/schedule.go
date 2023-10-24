package model

import "time"

type Schedule struct {
	Id        int       `gorm:"primary_key" json:"id"`
	PetId     int       `json:"pet_id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
