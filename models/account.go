package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/asaskevich/govalidator"
)

const (
	AccountTableName = "accounts"
)

type Account struct {
	gorm.Model
	Email    string `gorm:"unique_index" valid:"email,required"`
	Password string `valid:"length(8|32),matches(^[a-zA-Z0-9]+$),required"`
}

// hashPassword hashes the password for the given user
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (a *Account) Validate() error {
	if _, err := govalidator.ValidateStruct(a); err != nil {
		return err
	}
	return nil
}
