package model

import "time"

// User struct -  basic user model
type Pet struct {
	Id           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Age          string    `json:"age"`
	Sex          string    `json:"sex"`
	NowWeight    string    `json:"now_weight"`
	TargetWeight string    `json:"target_weight"`
	Birthday     string    `json:"birthday"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
