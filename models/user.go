package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Reward_id *int   `json:"reward_id"`
	Reward    Reward `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
