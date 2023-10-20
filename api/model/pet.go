package model

import "time"

// User struct -  basic user model
type Pet struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	BirthDay  time.Time `json:"birth_day"`
	Sex       int       `json:"sex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
