package models

import "github.com/jinzhu/gorm"

type RoleDS interface {
	Create(*Role) error
	Update(*Role) error
	Get(*Role) error
	GetList(uint) ([]Role, error)
	Delete(*Role) error
}

func RoleStore(db *gorm.DB) RoleDS {
	return &RoleSQL{
		db: db,
	}
}
