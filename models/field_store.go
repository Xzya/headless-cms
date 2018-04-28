package models

import "github.com/jinzhu/gorm"

type FieldDS interface {
	Create(*Field) error
	Update(*Field) error
	GetList(itemTypeID uint, projectID uint) ([]Field, error)
	GetListForProject(projectID uint, filter FilterOptions) ([]Field, error)
	Get(*Field) error
	Delete(*Field) error
}

func FieldStore(db *gorm.DB) FieldDS {
	return &FieldSQL{
		db: db,
	}
}
