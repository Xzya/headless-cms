package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"strings"
	"google.golang.org/grpc/status"
	apimodels "github.com/Xzya/headless-cms/gen/models"
)

type UserSQL struct {
	db *gorm.DB
}

func (s *UserSQL) Create(item *User) error {
	if item == nil || item.ProjectID == 0 || item.RoleID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	tx := s.db.Begin()

	// check if the user exists
	exists := !tx.Table(UserTableName).
		Where("LOWER(email) = ? AND project_id = ?", strings.ToLower(item.Email), item.ProjectID).
		Find(&User{}).
		Limit(1).RecordNotFound()
	if exists {
		tx.Rollback()
		return status.Errorf(codes.FailedPrecondition, "user already exists")
	}

	role := Role{
		Model: gorm.Model{
			ID: item.RoleID,
		},
	}

	if err := tx.Table(RoleTableName).Where("id = ? AND project_id = ?", item.RoleID, item.ProjectID).Find(&role).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "role not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	// check if there is already an account with this email
	var existingAccount Account

	if err := tx.Table(AccountTableName).Where("email = ?", item.Email).Find(&existingAccount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// no account
			item.State = apimodels.UserAttributesStateInvitationPending
		} else {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else if existingAccount.ID != 0 {
		// already registered
		item.State = apimodels.UserAttributesStateRegistered

		// make sure it's not the owner of the project
		var existingProject Project

		if err := tx.Table(ProjectTableName).Where("id = ?", item.ProjectID).Find(&existingProject).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		if existingAccount.ID == existingProject.AccountID {
			return status.Errorf(codes.FailedPrecondition, "the given user is the owner of the project")
		}

	} else {
		item.State = apimodels.UserAttributesStateInvitationPending
	}

	if err := tx.Table(UserTableName).Create(item).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	item.Role = role

	return nil
}

func (s *UserSQL) Update(item *User) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 || item.RoleID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	var existingItem User

	alreadyExists := !s.db.Table(UserTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {
		role := Role{
			Model: gorm.Model{
				ID: item.RoleID,
			},
		}

		if err := s.db.Table(RoleTableName).Where("id = ? AND project_id = ?", item.RoleID, item.ProjectID).Find(&role).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return status.Errorf(codes.InvalidArgument, "role not found")
			}
			return status.Errorf(codes.Internal, err.Error())
		}

		existingItem.RoleID = role.ID
		existingItem.Role = role

		if err := s.db.Table(UserTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Save(&existingItem).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "user not found")
	}

	*item = existingItem

	return nil
}

func (s *UserSQL) Get(item *User) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user")
	}

	var result User
	if err := s.db.Table(UserTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "user not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	*item = result

	return nil
}

func (s *UserSQL) GetList(projectID uint) ([]User, error) {
	if projectID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid project id")
	}

	items := make([]User, 0)

	if err := s.db.Table(UserTableName).Where("project_id = ?", projectID).Find(&items).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return items, nil
}

func (s *UserSQL) Delete(item *User) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid user")
	}

	var existingItem User

	alreadyExists := !s.db.Table(UserTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {
		if err := s.db.Table(UserTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Delete(Field{}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "user not found")
	}

	*item = existingItem

	return nil
}
