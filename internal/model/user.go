package model

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email    string `gorm:"size:255,not null,uniqueIndex" json:"email" validate:"required,email"`
	Username string `gorm:"size:255,not null,unique" json:"username" validate:"required"`
	Password string `gorm:"size:255,not null" json:"password" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	err := validator.New().Struct(u)
	
	if err != nil {
		return err
	}
	u.Password = hashPassword(u.Password)
	
	return err
}

func hashPassword(p string) string {
	salt := 10
	password := []byte(p)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashedPassword)
}
