package models

import "github.com/jinzhu/gorm"

type AccountDS interface {
	Create(*Account) error
	Authenticate(*Account) error
}

func AccountStore(db *gorm.DB) AccountDS {
	return &AccountSQL{
		db: db,
	}
}
