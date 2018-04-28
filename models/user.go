package models

import (
	"github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
)

const (
	UserTableName = "users"
)

type User struct {
	gorm.Model
	ProjectID uint
	RoleID    uint
	Email     string `valid:"email,required"`
	State     string

	Role Role
}

func (u *User) Validate() error {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}
	return nil
}
