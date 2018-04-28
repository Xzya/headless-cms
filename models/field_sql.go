package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"strings"
	"google.golang.org/grpc/status"
)

type FieldSQL struct {
	db *gorm.DB
}

func (s *FieldSQL) Create(item *Field) error {
	if item == nil || item.ProjectID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid field")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	exists := !s.db.Table(FieldTableName).
		Where("(label = ? AND project_id = ?) OR (LOWER(api_key) = ? AND project_id = ?)", item.Label, item.ProjectID, strings.ToLower(item.ApiKey), item.ProjectID).
		Find(&ItemType{}).
		Limit(1).RecordNotFound()
	if exists {
		return status.Errorf(codes.FailedPrecondition, "api_key or label already exists")
	}

	if err := s.db.Table(FieldTableName).Create(item).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	FindFieldRelationships(s.db, item)

	return nil
}

func (s *FieldSQL) Update(item *Field) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid field")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	var existingItem Field

	alreadyExists := !s.db.Table(FieldTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {

		item.ApiKey = strings.ToLower(item.ApiKey)
		existingItem.ApiKey = strings.ToLower(existingItem.ApiKey)

		// if the api key or name changed
		if item.ApiKey != existingItem.ApiKey || item.Label != existingItem.Label {
			// make sure the changed key does not already exist
			apiKeyAlreadyExists := !s.db.Table(FieldTableName).Where("id != ? AND ((label = ? AND project_id = ?) OR (LOWER(api_key) = ? AND project_id = ?))", item.ID, item.Label, item.ProjectID, item.ApiKey, item.ProjectID).
				Limit(1).
				Find(&ItemType{}).
				RecordNotFound()

			if apiKeyAlreadyExists {
				return status.Errorf(codes.FailedPrecondition, "api_key and name must be unique")
			}
		}

		existingItem.Label = item.Label
		existingItem.ApiKey = item.ApiKey
		existingItem.FieldType = item.FieldType
		existingItem.Localized = item.Localized
		existingItem.Hint = item.Hint
		existingItem.AppearanceType = item.AppearanceType
		existingItem.Position = item.Position
		existingItem.Required = item.Required
		existingItem.Unique = item.Unique
		existingItem.LengthMin = item.LengthMin
		existingItem.LengthMax = item.LengthMax
		existingItem.LengthEq = item.LengthEq
		existingItem.FormatCustomPattern = item.FormatCustomPattern
		existingItem.FormatPredefinedPattern = item.FormatPredefinedPattern
		existingItem.EnumValues = item.EnumValues

		if err := s.db.Table(FieldTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Save(&existingItem).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "field not found")
	}

	if err := FindFieldRelationships(s.db, &existingItem); err != nil {
		return err
	}

	*item = existingItem

	return nil
}

func (s *FieldSQL) GetList(itemTypeID uint, projectID uint) ([]Field, error) {
	if itemTypeID == 0 || projectID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid item type or project id")
	}

	items := make([]Field, 0)

	if err := s.db.Table(FieldTableName).Where("item_type_id = ? AND project_id = ?", itemTypeID, projectID).Find(&items).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for i := range items {
		FindFieldRelationships(s.db, &items[i])
	}

	return items, nil
}

func (s *FieldSQL) GetListForProject(projectID uint, filter FilterOptions) ([]Field, error) {
	if projectID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid project id")
	}

	items := make([]Field, 0)

	query := s.db.Table(FieldTableName).Where("project_id = ?", projectID)

	if len(filter.IDs) > 0 {
		// TODO: - test this, make sure only items for this project are returned
		query = query.Where("id IN (?)", filter.IDs)
	}
	if filter.TargetID > 0 && (filter.Target == "itemTypeId") {
		query = query.Where("item_type_id = ?", filter.TargetID)
	}

	if err := query.Find(&items).Error; err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for i := range items {
		FindFieldRelationships(s.db, &items[i])
	}

	return items, nil
}

func (s *FieldSQL) Get(item *Field) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid field")
	}

	var result Field
	if err := s.db.Table(FieldTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "field not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	FindFieldRelationships(s.db, &result)

	*item = result

	return nil
}

func (s *FieldSQL) Delete(item *Field) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid field")
	}

	var existingItem Field

	alreadyExists := !s.db.Table(FieldTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).
		Limit(1).
		Find(&existingItem).
		RecordNotFound()

	if alreadyExists {
		// TODO: - also delete strings etc. related to this field
		if err := s.db.Table(FieldTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Delete(Field{}).Error; err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
	} else {
		return status.Errorf(codes.NotFound, "field not found")
	}

	FindFieldRelationships(s.db, &existingItem)

	*item = existingItem

	return nil
}

func FindFieldRelationships(db *gorm.DB, item *Field) error {
	item.ItemType.ID = item.ItemTypeID
	return nil
}
