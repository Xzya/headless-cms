package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

type ProjectSQL struct {
	db *gorm.DB
}

func (s *ProjectSQL) Create(project *Project) error {
	if project == nil || project.AccountID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid project")
	}

	if err := project.Validate(); err != nil {
		return err
	}

	tx := s.db.Begin()

	if err := tx.Create(project).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	admin := DefaultAdminRole
	editor := DefaultEditorRole

	admin.ProjectID = project.ID
	editor.ProjectID = project.ID

	if err := tx.Table(RoleTableName).Create(&admin).Create(&editor).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (s *ProjectSQL) Update(project *Project) error {
	if project == nil || project.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid project")
	}

	if err := project.Validate(); err != nil {
		return err
	}

	var existingProject Project
	if err := s.db.Table(ProjectTableName).Where("id = ?", project.ID).Find(&existingProject).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "project not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	existingProject.Name = project.Name

	if err := s.db.Table(ProjectTableName).Save(&existingProject).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	*project = existingProject

	return nil
}

func (s *ProjectSQL) Get(project *Project) error {
	if project == nil || project.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid project")
	}

	var result Project
	if err := s.db.Table(ProjectTableName).Where("id = ?", project.ID).Find(&result).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	*project = result

	return nil
}

func (s *ProjectSQL) GetList() ([]Project, error) {
	projects := make([]Project, 0)

	if err := s.db.Table(ProjectTableName).Find(&projects).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return projects, nil
}

func (s *ProjectSQL) Delete(project *Project) error {
	if project == nil || project.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid project")
	}

	exists := !s.db.Table(ProjectTableName).Where("id = ?", project.ID).Find(project).RecordNotFound()
	if exists {
		// TODO: - also delete item-types, items, fields, strings etc. related to this project
		if err := s.db.Table(ProjectTableName).Where("id = ?", project.ID).Delete(Project{}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "project not found")
	}

	return nil
}
