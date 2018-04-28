package cms

import (
	"github.com/jinzhu/gorm"
	"github.com/Xzya/headless-cms/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/Xzya/headless-cms/config"
	"time"
	"github.com/Xzya/headless-cms/models"
)

var (
	JWTTokenExpiryTime time.Duration = 7 * 24 * time.Hour
)

type Service interface {
	// Account
	CreateAccount(*models.Account) (string, error)
	AuthenticateAccount(*models.Account) (string, error)

	// Project
	CreateProject(*models.Project) error
	UpdateProject(*models.Project) error
	GetProject(*models.Project) error
	GetProjectList() ([]models.Project, error)
	DeleteProject(*models.Project) error

	// Role
	CreateRole(*models.Role) error
	UpdateRole(*models.Role) error
	GetRole(*models.Role) error
	GetRoleList(uint) ([]models.Role, error)
	DeleteRole(*models.Role) error

	// User
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	GetUser(*models.User) error
	GetUserList(uint) ([]models.User, error)
	DeleteUser(*models.User) error

	// Item Type
	CreateItemType(*models.ItemType) error
	UpdateItemType(*models.ItemType) error
	GetItemTypeList(uint) ([]models.ItemType, error)
	GetItemType(*models.ItemType) error
	DuplicateItemType(*models.ItemType) error
	DeleteItemType(*models.ItemType) error

	// Field
	CreateField(*models.Field) error
	UpdateField(*models.Field) error
	GetFieldList(itemTypeID uint, projectID uint) ([]models.Field, error)
	GetProjectFieldList(projectID uint, filter models.FilterOptions) ([]models.Field, error)
	GetField(*models.Field) error
	DeleteField(*models.Field) error

	// Item
	CreateItem(*models.Item) error
	UpdateItem(*models.Item) error
	GetItemList(projectID uint, filter models.ItemFilterOptions) (int, []models.Item, error)
	GetItem(*models.Item) error
	DeleteItem(*models.Item) error
}

func NewService(db *gorm.DB, config config.Config) Service {
	return &service{
		Config:     config,
		AccountDS:  models.AccountStore(db),
		ProjectDS:  models.ProjectStore(db),
		RoleDS:     models.RoleStore(db),
		UserDS:     models.UserStore(db),
		ItemTypeDS: models.ItemTypeStore(db),
		FieldDS:    models.FieldStore(db),
		ItemDS:     models.ItemStore(db),
	}
}

type service struct {
	Config config.Config
	models.AccountDS
	models.ProjectDS
	models.RoleDS
	models.UserDS
	models.ItemTypeDS
	models.FieldDS
	models.ItemDS
}

// Account

func (s *service) CreateAccount(acc *models.Account) (string, error) {
	if err := s.AccountDS.Create(acc); err != nil {
		return "", err
	}

	token, err := auth.CreateToken(jwt.MapClaims{
		"sub":   acc.ID,
		"email": acc.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(JWTTokenExpiryTime).Unix(),
	}, s.Config.Secret)

	return token, err
}

func (s *service) AuthenticateAccount(acc *models.Account) (string, error) {
	if err := s.AccountDS.Authenticate(acc); err != nil {
		return "", err
	}

	token, err := auth.CreateToken(jwt.MapClaims{
		"sub":   acc.ID,
		"email": acc.Email,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(JWTTokenExpiryTime).Unix(),
	}, s.Config.Secret)

	return token, err
}

// Project
func (s *service) CreateProject(project *models.Project) error {
	return s.ProjectDS.Create(project)
}

func (s *service) UpdateProject(project *models.Project) error {
	return s.ProjectDS.Update(project)
}

func (s *service) GetProject(project *models.Project) error {
	return s.ProjectDS.Get(project)
}

func (s *service) GetProjectList() ([]models.Project, error) {
	return s.ProjectDS.GetList()
}

func (s *service) DeleteProject(project *models.Project) error {
	return s.ProjectDS.Delete(project)
}

// Role
func (s *service) CreateRole(role *models.Role) error {
	return s.RoleDS.Create(role)
}

func (s *service) UpdateRole(role *models.Role) error {
	return s.RoleDS.Update(role)
}

func (s *service) GetRole(role *models.Role) error {
	return s.RoleDS.Get(role)
}

func (s *service) GetRoleList(projectID uint) ([]models.Role, error) {
	return s.RoleDS.GetList(projectID)
}

func (s *service) DeleteRole(role *models.Role) error {
	return s.RoleDS.Delete(role)
}

// User
func (s *service) CreateUser(user *models.User) error {
	return s.UserDS.Create(user)
}

func (s *service) UpdateUser(user *models.User) error {
	return s.UserDS.Update(user)
}

func (s *service) GetUser(user *models.User) error {
	return s.UserDS.Get(user)
}

func (s *service) GetUserList(projectID uint) ([]models.User, error) {
	return s.UserDS.GetList(projectID)
}

func (s *service) DeleteUser(user *models.User) error {
	return s.UserDS.Delete(user)
}

// Item Type
func (s *service) CreateItemType(item *models.ItemType) error {
	return s.ItemTypeDS.Create(item)
}

func (s *service) UpdateItemType(item *models.ItemType) error {
	return s.ItemTypeDS.Update(item)
}

func (s *service) GetItemTypeList(projectID uint) ([]models.ItemType, error) {
	return s.ItemTypeDS.GetList(projectID)
}

func (s *service) GetItemType(item *models.ItemType) error {
	return s.ItemTypeDS.Get(item)
}

func (s *service) DuplicateItemType(item *models.ItemType) error {
	return s.ItemTypeDS.Duplicate(item)
}

func (s *service) DeleteItemType(item *models.ItemType) error {
	return s.ItemTypeDS.Delete(item)
}

// Field
func (s *service) CreateField(item *models.Field) error {
	return s.FieldDS.Create(item)
}

func (s *service) UpdateField(item *models.Field) error {
	return s.FieldDS.Update(item)
}

func (s *service) GetFieldList(itemTypeID uint, projectID uint) ([]models.Field, error) {
	return s.FieldDS.GetList(itemTypeID, projectID)
}

func (s *service) GetProjectFieldList(projectID uint, filter models.FilterOptions) ([]models.Field, error) {
	return s.FieldDS.GetListForProject(projectID, filter)
}

func (s *service) GetField(item *models.Field) error {
	return s.FieldDS.Get(item)
}

func (s *service) DeleteField(item *models.Field) error {
	return s.FieldDS.Delete(item)
}

// Item
func (s *service) CreateItem(item *models.Item) error {
	return s.ItemDS.Create(item)
}

func (s *service) UpdateItem(item *models.Item) error {
	return s.ItemDS.Update(item)
}

func (s *service) GetItemList(projectID uint, filter models.ItemFilterOptions) (int, []models.Item, error) {
	return s.ItemDS.GetList(projectID, filter)
}

func (s *service) GetItem(item *models.Item) error {
	return s.ItemDS.Get(item)
}

func (s *service) DeleteItem(item *models.Item) error {
	return s.ItemDS.Delete(item)
}
