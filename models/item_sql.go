package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"fmt"
	"time"
)

type ItemSQL struct {
	db *gorm.DB
}

func (s *ItemSQL) Create(item *Item) error {
	if item == nil || item.ProjectID == 0 || item.ItemTypeID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	tx := s.db.Begin()

	if err := tx.Table(ItemTableName).Create(item).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := CreateItemValues(tx, item); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	item.Attributes["updatedAt"] = time.Now().Format(time.RFC3339)

	return nil
}

func (s *ItemSQL) Update(item *Item) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item")
	}

	if err := item.Validate(); err != nil {
		return err
	}

	attributes := item.Attributes
	if err := s.Get(item); err != nil {
		return err
	}
	item.Attributes = attributes

	tx := s.db.Begin()

	if err := DeleteItemValues(tx, item); err != nil {
		tx.Rollback()
		return err
	}

	if err := CreateItemValues(tx, item); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(item).Updates(Item{Model: gorm.Model{UpdatedAt: time.Now()}}).Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	item.Attributes["updatedAt"] = time.Now().Format(time.RFC3339)

	return nil
}

func (s *ItemSQL) GetList(projectID uint, filter ItemFilterOptions) (int, []Item, error) {
	if projectID == 0 {
		return 0, nil, status.Errorf(codes.InvalidArgument, "invalid projectID")
	}

	count := 0
	items := make([]Item, 0)

	query := s.db.Table(ItemTableName).Select("*, COUNT(*) OVER()").Where("project_id = ?", projectID)

	if len(filter.IDs) > 0 {
		query = query.Where("id IN (?)", filter.IDs)
	}
	if filter.ItemTypeID > 0 {
		query = query.Where("item_type_id = ?", filter.ItemTypeID)
	}

	query = query.Count(&count)

	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	if err := query.Find(&items).Error; err != nil {
		return 0, nil, status.Errorf(codes.Internal, err.Error())
	}

	for i := range items {
		if err := s.Get(&items[i]); err != nil {
			return 0, nil, err
		}
	}

	return count, items, nil
}

func (s *ItemSQL) Get(item *Item) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item")
	}

	var result Item
	if err := s.db.Table(ItemTableName).Where("id = ? AND project_id = ?", item.ID, item.ProjectID).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "item not found")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := FindItemRelationships(s.db, &result); err != nil {
		return err
	}

	attributes := make(map[string]interface{})

	for _, field := range result.ItemType.Fields {
		switch field.FieldType {
		case apimodels.FieldAttributesFieldTypeString:

			str := String{}
			if err := s.db.Table(StringTableName).Where("item_id = ? AND field_id = ?", result.ID, field.ID).Find(&str).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					continue
				}
				return status.Errorf(codes.Internal, err.Error())
			}
			attributes[field.ApiKey] = str.Value

		default:
			// TODO: - support more types
		}
	}

	result.Attributes = attributes

	*item = result

	item.Attributes["updatedAt"] = result.UpdatedAt.Format(time.RFC3339)

	return nil
}

func (s *ItemSQL) Delete(item *Item) error {
	if item == nil || item.ProjectID == 0 || item.ID == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid item")
	}

	if err := s.Get(item); err != nil {
		return err
	}

	tx := s.db.Begin()

	if err := DeleteItemValues(tx, item); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(ItemTableName).Where("id = ?", item.ID).Delete(&Item{}).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	item.Attributes["updatedAt"] = item.UpdatedAt.Format(time.RFC3339)

	return nil
}

func FindItemRelationships(db *gorm.DB, item *Item) error {
	db.Model(item).Association("ItemType").Find(&item.ItemType)

	return FindItemTypeRelationships(db, &item.ItemType)
}

func CreateItemValues(db *gorm.DB, item *Item) error {
	if err := FindItemRelationships(db, item); err != nil {
		return err
	}

	attributes := make(map[string]interface{})

	for _, field := range item.ItemType.Fields {
		v, ok := item.Attributes[field.ApiKey]
		if ok {

			attributes[field.ApiKey] = v

			switch field.FieldType {
			case apimodels.FieldAttributesFieldTypeString:

				if str, ok := v.(string); ok {
					if err := ValidateString(db, item, &field, str); err != nil {
						return err
					}

					value := String{
						FieldID: field.ID,
						ItemID:  item.ID,
						Value:   str,
					}
					if err := db.Table(StringTableName).Create(&value).Error; err != nil {
						return status.Errorf(codes.Internal, err.Error())
					}
				} else {
					return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value for field `%v` must be a string", field.ApiKey))
				}

			default:
				// TODO: - add support for more types
			}
		} else if field.Required {
			return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("missing required field `%v`", field.ApiKey))
		}
	}

	item.Attributes = attributes

	return nil
}

func DeleteItemValues(db *gorm.DB, item *Item) error {
	if err := db.Table(StringTableName).Where("item_id = ?", item.ID).Delete(&String{}).Error; err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	// TODO: - remove any other type
	return nil
}
