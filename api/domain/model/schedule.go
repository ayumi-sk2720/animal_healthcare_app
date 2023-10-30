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
