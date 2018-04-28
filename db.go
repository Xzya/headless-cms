package cms

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Xzya/headless-cms/models"
)

// DatabaseConfig contains the database settings
type DatabaseConfig struct {
	Dialect string
	Source  string
	Debug   bool
}

func NewDatabase(config DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(config.Dialect, config.Source)
	if err != nil {
		return nil, err
	}

	if err := InitDatabase(db, config); err != nil {
		return nil, err
	}

	return db, nil
}

// InitDatabase initializes the database
func InitDatabase(db *gorm.DB, config DatabaseConfig) error {

	// set logging mode
	db.LogMode(config.Debug)

	//db.DropTableIfExists(Models()...)

	// auto migrate
	if err := db.AutoMigrate(models.Models()...).Error; err != nil {
		return err
	}

	// project belongs to an account
	db.Model(&models.Project{}).AddForeignKey("account_id", fmt.Sprintf("%s(id)", models.AccountTableName), "SET NULL", "CASCADE")

	// project can have many roles
	db.Model(&models.Role{}).AddForeignKey("project_id", fmt.Sprintf("%s(id)", models.ProjectTableName), "SET NULL", "CASCADE")

	// project can have many item types
	db.Model(&models.ItemType{}).AddForeignKey("project_id", fmt.Sprintf("%s(id)", models.ProjectTableName), "SET NULL", "CASCADE")

	// project can have many fields
	db.Model(&models.Field{}).AddForeignKey("project_id", fmt.Sprintf("%s(id)", models.ProjectTableName), "SET NULL", "CASCADE")

	// project can have many items
	db.Model(&models.Item{}).AddForeignKey("project_id", fmt.Sprintf("%s(id)", models.ProjectTableName), "SET NULL", "CASCADE")

	// project can have many roles
	db.Model(&models.Role{}).AddForeignKey("project_id", fmt.Sprintf("%s(id)", models.ProjectTableName), "SET NULL", "CASCADE")

	// project can have many users
	db.Model(&models.User{}).AddForeignKey("project_id", fmt.Sprintf("%s(id)", models.ProjectTableName), "SET NULL", "CASCADE")

	// user can have a role
	db.Model(&models.User{}).AddForeignKey("role_id", fmt.Sprintf("%s(id)", models.RoleTableName), "SET NULL", "CASCADE")

	// item type can have many fields
	db.Model(&models.ItemType{}).Related(&models.Field{}, "Fields")
	db.Model(&models.Field{}).Related(&models.ItemType{}, "ItemType")
	db.Model(&models.Field{}).AddForeignKey("item_type_id", fmt.Sprintf("%s(id)", models.ItemTypeTableName), "SET NULL", "CASCADE")

	// item type can have many items
	db.Model(&models.Item{}).Related(&models.ItemType{}, "ItemType")
	db.Model(&models.Item{}).AddForeignKey("item_type_id", fmt.Sprintf("%s(id)", models.ItemTypeTableName), "SET NULL", "CASCADE")

	// field can have many strings, etc.
	db.Model(&models.String{}).AddForeignKey("field_id", fmt.Sprintf("%s(id)", models.FieldTableName), "SET NULL", "CASCADE")

	// item can have many strings, etc.
	db.Model(&models.String{}).AddForeignKey("item_id", fmt.Sprintf("%s(id)", models.ItemTableName), "SET NULL", "CASCADE")

	// TODO: - add tests for these
	// item type: apiKey - project id pair is unique
	//db.Model(&models.ItemType{}).AddUniqueIndex("idx_api_key", "project_id", "api_key")
	// item type: name - project id pair is unique
	//db.Model(&models.ItemType{}).AddUniqueIndex("idx_name", "project_id", "name")

	// TODO: - add tests for these
	// field: apiKey - item type id pair is unique
	//db.Model(&models.Field{}).AddUniqueIndex("idx_api_key", "item_type_id", "api_key")
	// field: label - item type id pair is unique
	//db.Model(&models.Field{}).AddUniqueIndex("idx_label", "item_type_id", "label")

	return nil
}
