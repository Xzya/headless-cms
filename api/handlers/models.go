package handlers

import (
	"github.com/Xzya/headless-cms/models"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/strfmt"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

// Account

func AccountToRel(f *models.Account) *apimodels.Relationship {
	if f != nil {
		return &apimodels.Relationship{
			ID:   UintToString(f.ID),
			Type: apimodels.RelationshipTypeAccount,
		}
	}
	return nil
}

// Project

func ProjectToSwag(f *models.Project) *apimodels.Project {
	if f != nil {
		account := &apimodels.ProjectRelationshipsAccount{}

		if f.AccountID != 0 {
			account.Data = &apimodels.Relationship{
				ID:   UintToString(f.AccountID),
				Type: apimodels.RelationshipTypeAccount,
			}
		}

		return &apimodels.Project{
			ID:   UintToString(f.ID),
			Type: apimodels.ProjectTypeProject,
			Attributes: &apimodels.ProjectAttributes{
				Name: f.Name,
			},
			Relationships: &apimodels.ProjectRelationships{
				Account: account,
			},
		}
	}
	return nil
}

func SwagToProject(f *apimodels.Project) *models.Project {
	if f != nil {
		return &models.Project{
			Model: gorm.Model{
				ID: StringToUint(f.ID),
			},
			Name: f.Attributes.Name,
		}
	}
	return nil
}

func ProjectWithID(id uint) *models.Project {
	return &models.Project{
		Model: gorm.Model{
			ID: id,
		},
	}
}

// Role

func RoleToSwag(f *models.Role) *apimodels.Role {
	if f != nil {
		return &apimodels.Role{
			ID:   UintToString(f.ID),
			Type: apimodels.RoleTypeRole,
			Attributes: apimodels.RoleAttributes{
				Name:                  f.Name,
				CanEditProject:        f.CanEditProject,
				CanEditSchema:         f.CanEditSchema,
				CanManageUsers:        f.CanManageUsers,
				CanManageAccessTokens: f.CanManageAccessTokens,
				CanPublishContent:     f.CanPublishContent,
			},
		}
	}
	return nil
}

func SwagToRole(f *apimodels.Role) *models.Role {
	if f != nil {
		return &models.Role{
			Model: gorm.Model{
				ID: StringToUint(f.ID),
			},
			Name:                  f.Attributes.Name,
			CanEditProject:        f.Attributes.CanEditProject,
			CanEditSchema:         f.Attributes.CanEditSchema,
			CanManageUsers:        f.Attributes.CanManageUsers,
			CanManageAccessTokens: f.Attributes.CanManageAccessTokens,
			CanPublishContent:     f.Attributes.CanPublishContent,
		}
	}
	return nil
}

func RoleWithID(id uint) *models.Role {
	return &models.Role{
		Model: gorm.Model{
			ID: id,
		},
	}
}

// User

func UserToSwag(f *models.User) *apimodels.User {
	if f != nil {
		role := &apimodels.UserRelationshipsRole{}

		if f.RoleID != 0 {
			role.Data = &apimodels.Relationship{
				ID:   UintToString(f.RoleID),
				Type: apimodels.RoleTypeRole,
			}
		}

		return &apimodels.User{
			ID:   UintToString(f.ID),
			Type: apimodels.UserTypeUser,
			Attributes: &apimodels.UserAttributes{
				Email: strfmt.Email(f.Email),
				State: f.State,
			},
			Relationships: &apimodels.UserRelationships{
				Role: role,
			},
		}
	}
	return nil
}

func SwagToUser(f *apimodels.User) *models.User {
	if f != nil {
		return &models.User{
			Model: gorm.Model{
				ID: StringToUint(f.ID),
			},
			Email:  f.Attributes.Email.String(),
			State:  f.Attributes.State,
			RoleID: StringToUint(f.Relationships.Role.Data.ID),
		}
	}
	return nil
}

func UserWithID(id uint) *models.User {
	return &models.User{
		Model: gorm.Model{
			ID: id,
		},
	}
}

// Item Type

func ItemTypeToSwag(f *models.ItemType) *apimodels.ItemType {
	if f != nil {
		fields := &apimodels.ItemTypeRelationshipsFields{
			Data: make([]*apimodels.Relationship, 0),
		}
		orderingField := &apimodels.ItemTypeRelationshipsOrderingField{}
		singletonItem := &apimodels.ItemTypeRelationshipsSingletonItem{}

		if len(f.Fields) > 0 {
			for _, item := range f.Fields {
				newItem := FieldToRel(&item)
				fields.Data = append(fields.Data, newItem)
			}
		}

		return &apimodels.ItemType{
			ID:   UintToString(f.ID),
			Type: apimodels.ItemTypeTypeItemType,
			Attributes: &apimodels.ItemTypeAttributes{
				Name:   f.Name,
				APIKey: f.ApiKey,
			},
			Relationships: &apimodels.ItemTypeRelationships{
				Fields:        fields,
				OrderingField: orderingField,
				SingletonItem: singletonItem,
			},
		}
	}
	return nil
}

func SwagToItemType(f *apimodels.ItemType) *models.ItemType {
	if f != nil {
		return &models.ItemType{
			Model: gorm.Model{
				ID: StringToUint(f.ID),
			},
			Name:   f.Attributes.Name,
			ApiKey: f.Attributes.APIKey,
		}
	}
	return nil
}

func ItemTypeWithID(id uint) *models.ItemType {
	return &models.ItemType{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func ItemTypeIncluded(f *models.ItemType) apimodels.ItemTypeIncluded {
	included := make(apimodels.ItemTypeIncluded, 0)
	if f != nil {
		for _, field := range f.Fields {
			included = append(included, FieldToSwag(&field))
		}
	}
	return included
}

func ItemTypeToRel(f *models.ItemType) *apimodels.Relationship {
	if f != nil {
		return &apimodels.Relationship{
			ID:   UintToString(f.ID),
			Type: apimodels.ItemTypeTypeItemType,
		}
	}
	return nil
}

// Field

func FieldToSwag(f *models.Field) *apimodels.Field {
	if f != nil {
		// validators
		var unique interface{}
		var required interface{}
		var enum *apimodels.FieldAttributesValidatorsEnum
		var format *apimodels.FieldAttributesValidatorsFormat
		var length *apimodels.FieldAttributesValidatorsLength

		if f.Unique {
			unique = make(map[string]interface{})
		}
		if f.Required {
			required = make(map[string]interface{})
		}
		if len(f.EnumValues) > 0 {
			enum = &apimodels.FieldAttributesValidatorsEnum{
				Values: f.EnumValues,
			}
		}
		if f.FormatCustomPattern != nil || f.FormatPredefinedPattern != nil {
			format = &apimodels.FieldAttributesValidatorsFormat{
				CustomPattern:     StringValueIfNotEmpty(f.FormatCustomPattern),
				PredefinedPattern: StringValueIfNotEmpty(f.FormatPredefinedPattern),
			}
		}
		if f.LengthMin != nil || f.LengthMax != nil || f.LengthEq != nil {
			length = &apimodels.FieldAttributesValidatorsLength{
				Min: Int32PtrToStringPtr(f.LengthMin),
				Max: Int32PtrToStringPtr(f.LengthMax),
				Eq:  Int32PtrToStringPtr(f.LengthEq),
			}
		}

		// relationships
		itemType := &apimodels.FieldRelationshipsItemType{}

		if f.ItemType.ID != 0 {
			itemType.Data = ItemTypeToRel(&f.ItemType)
		}

		return &apimodels.Field{
			ID:   UintToString(f.ID),
			Type: apimodels.FieldTypeField,
			Attributes: &apimodels.FieldAttributes{
				Label:     f.Label,
				APIKey:    f.ApiKey,
				Position:  f.Position,
				Localized: f.Localized,
				Hint:      f.Hint,
				FieldType: f.FieldType,
				Appearance: apimodels.FieldAttributesAppearance{
					Type: f.AppearanceType,
				},
				Validators: apimodels.FieldAttributesValidators{
					Unique:   unique,
					Required: required,
					Enum:     enum,
					Format:   format,
					Length:   length,
				},
			},
			Relationships: &apimodels.FieldRelationships{
				ItemType: itemType,
			},
		}
	}
	return nil
}

func SwagToField(f *apimodels.Field) *models.Field {
	if f != nil {
		var unique bool
		var required bool
		var enum pq.StringArray
		var formatCustomPattern *string
		var formatPredefinedPattern *string
		var lengthMin *int32
		var lengthMax *int32
		var lengthEq *int32

		if f.Attributes.Validators.Unique != nil {
			unique = true
		}
		if f.Attributes.Validators.Required != nil {
			required = true
		}
		if f.Attributes.Validators.Enum != nil {
			enum = pq.StringArray(f.Attributes.Validators.Enum.Values)
		}
		if f.Attributes.Validators.Format != nil {
			formatCustomPattern = StringValueIfNotEmpty(f.Attributes.Validators.Format.CustomPattern)
			formatPredefinedPattern = StringValueIfNotEmpty(f.Attributes.Validators.Format.PredefinedPattern)
		}
		if f.Attributes.Validators.Length != nil {
			lengthMin = StringPtrToInt32Ptr(f.Attributes.Validators.Length.Min)
			lengthMax = StringPtrToInt32Ptr(f.Attributes.Validators.Length.Max)
			lengthEq = StringPtrToInt32Ptr(f.Attributes.Validators.Length.Eq)
		}

		return &models.Field{
			Model: gorm.Model{
				ID: StringToUint(f.ID),
			},
			Label:                   f.Attributes.Label,
			ApiKey:                  f.Attributes.APIKey,
			Position:                f.Attributes.Position,
			Localized:               f.Attributes.Localized,
			Hint:                    f.Attributes.Hint,
			FieldType:               f.Attributes.FieldType,
			AppearanceType:          f.Attributes.Appearance.Type,
			Unique:                  unique,
			Required:                required,
			EnumValues:              enum,
			FormatCustomPattern:     formatCustomPattern,
			FormatPredefinedPattern: formatPredefinedPattern,
			LengthMin:               lengthMin,
			LengthMax:               lengthMax,
			LengthEq:                lengthEq,
		}
	}
	return nil
}

func FieldWithID(id uint) *models.Field {
	return &models.Field{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func FieldToRel(f *models.Field) *apimodels.Relationship {
	if f != nil {
		return &apimodels.Relationship{
			ID:   UintToString(f.ID),
			Type: apimodels.FieldTypeField,
		}
	}
	return nil
}

func SwagToFieldFilter(params field.GetProjectFieldsParams) models.FilterOptions {
	ids := make([]uint, 0)
	for _, id := range params.Ids {
		ids = append(ids, uint(id))
	}
	return models.FilterOptions{
		IDs:      ids,
		Target:   swag.StringValue(params.Target),
		TargetID: uint(swag.Int32Value(params.ID)),
	}
}

// Item

func ItemToSwag(f *models.Item) *apimodels.Item {
	if f != nil {
		// relationships
		itemType := &apimodels.ItemRelationshipsItemType{}
		lastEditor := &apimodels.ItemRelationshipsLastEditor{}

		if f.ItemType.ID != 0 {
			itemType.Data = ItemTypeToRel(&f.ItemType)
		}

		return &apimodels.Item{
			ID:         UintToString(f.ID),
			Type:       apimodels.ItemTypeItem,
			Attributes: f.Attributes,
			Relationships: &apimodels.ItemRelationships{
				ItemType: itemType,
				// TODO: - add this
				LastEditor: lastEditor,
			},
		}
	}
	return nil
}

func SwagToItem(f *apimodels.Item) *models.Item {
	if f != nil {
		return &models.Item{
			Model: gorm.Model{
				ID: StringToUint(f.ID),
			},
			ItemTypeID: StringToUint(f.Relationships.ItemType.Data.ID),
			Attributes: f.Attributes,
		}
	}
	return nil
}

func ItemWithID(id uint) *models.Item {
	return &models.Item{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func SwagToItemFilter(params item.GetItemsParams) models.ItemFilterOptions {
	ids := make([]int, 0)
	for _, id := range params.Ids {
		ids = append(ids, int(id))
	}
	return models.ItemFilterOptions{
		ItemTypeID: uint(swag.Int32Value(params.Type)),
		Offset:     int(swag.Int32Value(params.Offset)),
		Limit:      int(swag.Int32Value(params.Limit)),
		IDs:        ids,
		Query:      swag.StringValue(params.Query),
	}
}
