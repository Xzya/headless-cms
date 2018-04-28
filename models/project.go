package models

import "github.com/jinzhu/gorm"

const (
	ProjectTableName = "projects"
)

type Project struct {
	gorm.Model
	AccountID uint
	Name      string
}

func (p *Project) Validate() error {
	return nil
}
