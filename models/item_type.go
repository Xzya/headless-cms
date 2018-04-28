package models

import "github.com/jinzhu/gorm"

const (
	ItemTypeTableName = "item_types"
)

type ItemType struct {
	gorm.Model
	ProjectID uint
	Name      string `gorm:"primary_key"`
	ApiKey    string `gorm:"primary_key"`

	Fields []Field
}

func (t ItemType) Validate() error {
	return nil
}
