package models

import (
	"time"

	"gorm.io/gorm"
)

type Reward struct {
	ID         int `json:"id" gorm:"primary_key"`
	Earn_point int `json:"earn_point"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
