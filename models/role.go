package models

import "github.com/jinzhu/gorm"

const (
	RoleTableName = "roles"
)

var (
	DefaultAdminRole = Role{
		Name:                  "Admin",
		CanEditProject:        true,
		CanEditSchema:         true,
		CanManageUsers:        true,
		CanManageAccessTokens: true,
		CanPublishContent:     true,
	}
	DefaultEditorRole = Role{
		Name:                  "Editor",
		CanEditProject:        false,
		CanEditSchema:         false,
		CanManageUsers:        false,
		CanManageAccessTokens: false,
		CanPublishContent:     true,
	}
)

type Role struct {
	gorm.Model
	ProjectID             uint
	Name                  string
	CanEditProject        bool
	CanEditSchema         bool
	CanManageUsers        bool
	CanManageAccessTokens bool
	CanPublishContent     bool
}

func (r *Role) Validate() error {
	return nil
}
