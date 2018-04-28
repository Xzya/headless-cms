package models

import "github.com/jinzhu/gorm"

type ProjectDS interface {
	Create(*Project) error
	Update(*Project) error
	Get(*Project) error
	GetList() ([]Project, error)
	Delete(*Project) error
}

func ProjectStore(db *gorm.DB) ProjectDS {
	return &ProjectSQL{
		db: db,
	}
}
