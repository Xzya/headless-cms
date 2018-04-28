package models

import (
	"github.com/jinzhu/gorm"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"regexp"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"fmt"
	"github.com/asaskevich/govalidator"
)

func ValidatePredefinedPattern(field *Field, v string) error {
	if pattern := field.FormatPredefinedPattern; pattern != nil {
		switch *pattern {
		case apimodels.FieldAttributesValidatorsFormatPredefinedPatternURL:
			if !govalidator.IsURL(v) {
				return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value `%v` for field `%v` is not a valid url", v, field.ApiKey))
			}
		case apimodels.FieldAttributesValidatorsFormatPredefinedPatternEmail:
			if !govalidator.IsEmail(v) {
				return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value `%v` for field `%v` is not a valid email", v, field.ApiKey))
			}
		default:
			// unknown pattern, ignore?
		}
	}
	return nil
}

func ValidateCustomPattern(field *Field, v string) error {
	if pattern := field.FormatCustomPattern; pattern != nil {
		regex, err := regexp.Compile(*pattern)
		if err != nil {
			return status.Errorf(codes.FailedPrecondition, err.Error())
		}
		if match := regex.MatchString(v); !match {
			return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value `%v` for field `%v` does not match `%v`", v, field.ApiKey, *pattern))
		}
	}
	return nil
}

func ValidateEnumValues(field *Field, v string) error {
	if len(field.EnumValues) > 0 {
		match := false
		for _, enumValue := range field.EnumValues {
			if enumValue == v {
				match = true
				break
			}
		}
		if !match {
			return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value `%v` must match one of the following enum values `%v`", v, field.EnumValues))
		}
	}
	return nil
}

func ValidateLength(field *Field, v string) error {
	if field.LengthMin != nil || field.LengthMax != nil || field.LengthEq != nil {
		if min := field.LengthMin; min != nil {
			if len(v) < int(*min) {
				return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value for field `%v` must be at least %v characters long", field.ApiKey, *min))
			}
		}
		if max := field.LengthMax; max != nil {
			if len(v) > int(*max) {
				return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value for field `%v` must be at most %v characters long", field.ApiKey, *max))
			}
		}
		if eq := field.LengthEq; eq != nil {
			if len(v) != int(*eq) {
				return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value for field `%v` must be exactly %v characters long", field.ApiKey, *eq))
			}
		}
	}
	return nil
}

func ValidateUniqueString(db *gorm.DB, field *Field, v string) error {
	if field.Unique {
		exists := !db.Table(StringTableName).Where("field_id = ? AND value = ?", field.ID, v).Limit(1).Find(&String{}).RecordNotFound()
		if exists {
			return status.Errorf(codes.FailedPrecondition, fmt.Sprintf("value for field `%v` must be unique", field.ApiKey))
		}
	}
	return nil
}

func ValidateString(db *gorm.DB, item *Item, field *Field, v string) error {
	if err := ValidatePredefinedPattern(field, v); err != nil {
		return err
	}

	if err := ValidateCustomPattern(field, v); err != nil {
		return err
	}

	if err := ValidateEnumValues(field, v); err != nil {
		return err
	}

	if err := ValidateLength(field, v); err != nil {
		return err
	}

	if err := ValidateUniqueString(db, field, v); err != nil {
		return err
	}

	return nil
}
