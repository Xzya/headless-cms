package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"fmt"
)

type RoleSQL struct {
	db *gorm.DB
}

func (s *RoleSQL) Create(item *Role) error {
	if item == nil || item.ProjectID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid role")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	tx := s.db.Begin()

	exists := !tx.Table(RoleTableName).
		Where("name = ? AND project_id = ?", item.Name, item.ProjectID).
		Find(&Role{}).
		Limit(1).RecordNotFound()
	if exists {
		tx.Rollback()
		return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("role with name `%v` already exists", item.Name))
	}

	if err := tx.Table(RoleTableName).Create(item).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (s *RoleSQL) Update(item *Role) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid role")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	var existingItem Role

	alreadyExists := !s.db.Table(RoleTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {

		existingItem.CanEditProject = item.CanEditProject
		existingItem.CanEditSchema = item.CanEditSchema
		existingItem.CanManageUsers = item.CanManageUsers
		existingItem.CanManageAccessTokens = item.CanManageAccessTokens
		existingItem.CanPublishContent = item.CanPublishContent

		if err := s.db.Table(RoleTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Save(&existingItem).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "role not found")
	}

	*item = existingItem

	return nil
}

func (s *RoleSQL) Get(item *Role) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.FailedPrecondition, "invalid role")
	}

	var result Role
	if err := s.db.Table(RoleTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.FailedPrecondition, "role not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	*item = result

	return nil
}

func (s *RoleSQL) GetList(projectID uint) ([]Role, error) {
	if projectID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid project id")
	}

	items := make([]Role, 0)

	if err := s.db.Table(RoleTableName).Where("project_id = ?", projectID).Find(&items).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return items, nil
}

func (s *RoleSQL) Delete(item *Role) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid role")
	}

	var existingItem Role

	alreadyExists := !s.db.Table(RoleTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {
		if err := s.db.Table(RoleTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Delete(Field{}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "role not found")
	}

	*item = existingItem

	return nil
}
