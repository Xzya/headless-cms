package models

import "github.com/jinzhu/gorm"

type UserDS interface {
	Create(*User) error
	Update(*User) error
	Get(*User) error
	GetList(uint) ([]User, error)
	Delete(*User) error
}

func UserStore(db *gorm.DB) UserDS {
	return &UserSQL{
		db: db,
	}
}
