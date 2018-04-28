package models

import (
	"github.com/jinzhu/gorm"
)

const (
	ItemTableName = "items"
)

type Item struct {
	gorm.Model
	ItemTypeID uint
	ProjectID  uint

	ItemType ItemType

	Attributes map[string]interface{} `sql:"-"`
}

func (f *Item) Validate() error {
	return nil
}

// TODO: - replace with `FilterOptions`

type ItemFilterOptions struct {
	IDs        []int
	ItemTypeID uint
	Offset     int
	Limit      int
	// TODO: - support this
	Query string
}
