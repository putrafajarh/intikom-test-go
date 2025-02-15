package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Email     string    `json:"email" gorm:"not null;unique;type:varchar(255)"`
	Password  string    `json:"-" gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Tasks     []Task    `json:"tasks,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) TableName() string {
	return "users"
}
