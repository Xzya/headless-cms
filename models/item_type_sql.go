package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"fmt"
	"strings"
)

type ItemTypeSQL struct {
	db *gorm.DB
}

func (s *ItemTypeSQL) Create(item *ItemType) error {
	if item == nil || item.ProjectID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item type")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	exists := !s.db.Table(ItemTypeTableName).
		Where("(name = ? AND project_id = ?) OR (LOWER(api_key) = ? AND project_id = ?)", item.Name, item.ProjectID, strings.ToLower(item.ApiKey), item.ProjectID).
		Find(&ItemType{}).
		Limit(1).RecordNotFound()
	if exists {
		return status.Errorf(codes.FailedPrecondition, "api_key or name already exists")
	}

	if err := s.db.Table(ItemTypeTableName).Create(item).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (s *ItemTypeSQL) Update(item *ItemType) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item type")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	var existingItem ItemType

	alreadyExists := !s.db.Table(ItemTypeTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {

		item.ApiKey = strings.ToLower(item.ApiKey)
		existingItem.ApiKey = strings.ToLower(existingItem.ApiKey)

		// if the api key or name changed
		if item.ApiKey != existingItem.ApiKey || item.Name != existingItem.Name {
			// make sure the changed key does not already exist
			apiKeyAlreadyExists := !s.db.Table(ItemTypeTableName).Where("id != ? AND ((name = ? AND project_id = ?) OR (LOWER(api_key) = ? AND project_id = ?))", item.ID, item.Name, item.ProjectID, item.ApiKey, item.ProjectID).
				Limit(1).
				Find(&ItemType{}).
				RecordNotFound()

			if apiKeyAlreadyExists {
				return status.Errorf(codes.FailedPrecondition, "api_key and name must be unique")
			}
		}

		existingItem.Name = item.Name
		existingItem.ApiKey = item.ApiKey

		if err := s.db.Table(ItemTypeTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Updates(ItemType{
			Name:   item.Name,
			ApiKey: item.ApiKey,
		}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "item type not found")
	}

	if err := FindItemTypeRelationships(s.db, &existingItem); err != nil {
		return err
	}

	*item = existingItem

	return nil
}

func (s *ItemTypeSQL) GetList(projectID uint) ([]ItemType, error) {
	if projectID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid projectID")
	}

	items := make([]ItemType, 0)

	if err := s.db.Table(ItemTypeTableName).Where("project_id = ?", projectID).Find(&items).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for i := range items {
		if err := FindItemTypeRelationships(s.db, &items[i]); err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (s *ItemTypeSQL) Get(item *ItemType) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item type")
	}

	var result ItemType
	if err := s.db.Table(ItemTypeTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "item type not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := FindItemTypeRelationships(s.db, &result); err != nil {
		return err
	}

	*item = result

	return nil
}

func (s *ItemTypeSQL) Duplicate(item *ItemType) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item type")
	}

	var result ItemType
	if err := s.db.Table(ItemTypeTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "item type not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	// TODO: - search for a better solution
	// we can get all items that start with the api_key
	// then compare them alphabetically
	// select the greatest one
	// if the last character is a number, increment
	// otherwise add _copy_1

	items := make([]ItemType, 0)
	if err := s.db.Table(ItemTypeTableName).Where("api_key LIKE LOWER(?) AND project_id = ?", fmt.Sprintf("%v%%", result.ApiKey), result.ProjectID).Find(&items).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	greatest := result.ApiKey
	for _, item := range items {
		if strings.Compare(item.ApiKey, greatest) == 1 {
			greatest = item.ApiKey
		}
	}

	result.ID = 0
	result.ApiKey = fmt.Sprintf("%s_copy", greatest)

	if err := s.db.Table(ItemTypeTableName).Create(&result).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := FindItemTypeRelationships(s.db, &result); err != nil {
		return err
	}

	*item = result

	return nil
}

func (s *ItemTypeSQL) Delete(item *ItemType) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item type")
	}

	var existingItem ItemType

	alreadyExists := !s.db.Table(ItemTypeTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {
		// TODO: - also delete items, fields, strings etc. related to this item
		if err := s.db.Table(ItemTypeTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Delete(ItemType{}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "item type not found")
	}

	if err := FindItemTypeRelationships(s.db, &existingItem); err != nil {
		return err
	}

	*item = existingItem

	return nil
}

func FindItemTypeRelationships(db *gorm.DB, item *ItemType) error {
	db.Model(item).Association("Fields").Find(&item.Fields)

	for i := range item.Fields {
		if err := FindFieldRelationships(db, &item.Fields[i]); err != nil {
			return err
		}
	}

	return nil
}
