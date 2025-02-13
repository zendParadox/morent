package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Struct untuk tabel users
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}

// Method untuk memeriksa password hash
func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}