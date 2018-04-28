package models

import "github.com/jinzhu/gorm"

type ItemTypeDS interface {
	Create(*ItemType) error
	Update(*ItemType) error
	GetList(uint) ([]ItemType, error)
	Get(*ItemType) error
	Duplicate(*ItemType) error
	Delete(*ItemType) error
}

func ItemTypeStore(db *gorm.DB) ItemTypeDS {
	return &ItemTypeSQL{
		db: db,
	}
}
