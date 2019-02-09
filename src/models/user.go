package models

import (
	"time"
)

type User struct {
	ID           uint       `gorm:"primary_key"`
	Name         string     `gorm:"column:name" json:"name"`
	Email        string     `gorm:"column:email" json:"email"`
	ImageProfile string     `gorm:"column:image_profile" json:"image_profile"`
	CreatedAt    *time.Time `sql:"index"`
	UpdatedAt    *time.Time `sql:"index"`
	DeletedAt    *time.Time `sql:"index"`
}

func (User) TableName() string {
	return "rl_users"
}
