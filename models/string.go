package models

import "github.com/jinzhu/gorm"

const (
	StringTableName = "strings"
)

type String struct {
	gorm.Model
	ItemID  uint
	FieldID uint

	Value string
}

func (f *String) Validate() error {
	return nil
}
