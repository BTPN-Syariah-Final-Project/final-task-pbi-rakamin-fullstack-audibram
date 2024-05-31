package models

import (
	"btpn-final/helpers"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"not null"`
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Photos    []Photo `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPassword(u.Password)
	return
}
