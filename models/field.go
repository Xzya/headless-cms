package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"fmt"
	"regexp"
)

const (
	FieldTableName = "fields"
)

var (
	ReservedKeys = []string{"updatedAt", "isValid"}
)

type Field struct {
	gorm.Model
	ItemTypeID     uint
	ProjectID      uint
	Label          string `gorm:"primary_key"`
	ApiKey         string `gorm:"primary_key"`
	FieldType      string
	Localized      bool
	Hint           string
	AppearanceType string
	Position       int32

	Required                bool
	Unique                  bool
	LengthMin               *int32
	LengthMax               *int32
	LengthEq                *int32
	FormatCustomPattern     *string
	FormatPredefinedPattern *string
	EnumValues              pq.StringArray `gorm:"type:text[]"`

	ItemType ItemType
}

func (f *Field) Validate() error {
	switch f.FieldType {
	case apimodels.FieldAttributesFieldTypeString:
		if f.FormatPredefinedPattern != nil && f.FormatCustomPattern != nil {
			return status.Errorf(codes.FailedPrecondition, "field cannot have both `predefinedPattern` and `customPattern` set")
		}
		if pattern := f.FormatPredefinedPattern; pattern != nil {
			if *pattern != apimodels.FieldAttributesValidatorsFormatPredefinedPatternURL &&
				*pattern != apimodels.FieldAttributesValidatorsFormatPredefinedPatternEmail {
				return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("unknown predefined pattern `%v`", *pattern))
			}
		}
		if pattern := f.FormatCustomPattern; pattern != nil {
			if _, err := regexp.Compile(*pattern); err != nil {
				return status.Errorf(codes.FailedPrecondition, err.Error())
			}
		}
		if f.EnumValues != nil && len(f.EnumValues) == 0 {
			return status.Errorf(codes.FailedPrecondition, "enum must have at least one value")
		}
		if (f.LengthMin != nil && f.LengthMax != nil) && f.LengthEq != nil {
			return status.Errorf(codes.FailedPrecondition, "length must not have both `eq` and `min` and/or `max`")
		}
		if f.AppearanceType != apimodels.FieldAttributesAppearanceTypeTitle &&
			f.AppearanceType != apimodels.FieldAttributesAppearanceTypePlain {
			return status.Errorf(codes.FailedPrecondition, "appearance type must be `title` or `plain`")
		}

		// TODO: - set all fields required for other types to nil
	}

	for _, reservedKey := range ReservedKeys {
		if f.ApiKey == reservedKey {
			return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("key `%v` is reserved", reservedKey))
		}
	}

	return nil
}
