package models

import "github.com/jinzhu/gorm"

type ItemDS interface {
	Create(*Item) error
	Update(*Item) error
	GetList(projectID uint, filter ItemFilterOptions) (int, []Item, error)
	Get(*Item) error
	Delete(*Item) error
}

func ItemStore(db *gorm.DB) ItemDS {
	return &ItemSQL{
		db: db,
	}
}
