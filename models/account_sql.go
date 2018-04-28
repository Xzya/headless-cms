package models

import (
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"golang.org/x/crypto/bcrypt"
	"strings"
	apimodels "github.com/Xzya/headless-cms/gen/models"
)

type AccountSQL struct {
	db *gorm.DB
}

func (s *AccountSQL) Create(acc *Account) error {
	if acc == nil {
		return status.Errorf(codes.InvalidArgument, "nil account")
	}

	if err := acc.Validate(); err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	if alreadyExists := !s.db.Table(AccountTableName).Where("LOWER(email) = ?", strings.ToLower(acc.Email)).Find(new(Account)).RecordNotFound(); alreadyExists {
		return status.Errorf(codes.AlreadyExists, "account already exists")
	}

	hash, err := hashPassword(acc.Password)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	acc.Password = hash

	acc.Email = strings.ToLower(acc.Email)

	tx := s.db.Begin()

	if err := tx.Create(acc).Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	// update the state of any existing users with this email
	if err := tx.Table(UserTableName).Where("LOWER(email) = ?", acc.Email).Updates(User{State: apimodels.UserAttributesStateRegistered}).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			tx.Rollback()
			return status.Errorf(codes.Internal, err.Error())
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}

func (s *AccountSQL) Authenticate(acc *Account) error {
	if acc == nil {
		return status.Errorf(codes.InvalidArgument, "nil account")
	}

	if len(acc.Email) == 0 || len(acc.Password) == 0 {
		return status.Errorf(codes.InvalidArgument, "missing email or password")
	}

	givenPassword := acc.Password

	var result Account

	if err := s.db.Table(AccountTableName).Where("LOWER(email) = ?", strings.ToLower(acc.Email)).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Errorf(codes.InvalidArgument, "invalid email or password")
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(givenPassword)); err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid email or password")
	}

	*acc = result
	acc.Password = ""

	return nil
}
