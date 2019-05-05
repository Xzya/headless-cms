// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Xzya/headless-cms/gen/restapi/operations/account"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

// NewCmsAPI creates a new Cms instance
func NewCmsAPI(spec *loads.Document) *CmsAPI {
	return &CmsAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		FieldCreateFieldHandler: field.CreateFieldHandlerFunc(func(params field.CreateFieldParams) middleware.Responder {
			return middleware.NotImplemented("operation FieldCreateField has not yet been implemented")
		}),
		ItemCreateItemHandler: item.CreateItemHandlerFunc(func(params item.CreateItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemCreateItem has not yet been implemented")
		}),
		ItemTypeCreateItemTypeHandler: item_type.CreateItemTypeHandlerFunc(func(params item_type.CreateItemTypeParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemTypeCreateItemType has not yet been implemented")
		}),
		ProjectCreateProjectHandler: project.CreateProjectHandlerFunc(func(params project.CreateProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation ProjectCreateProject has not yet been implemented")
		}),
		RoleCreateRoleHandler: role.CreateRoleHandlerFunc(func(params role.CreateRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation RoleCreateRole has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserCreateUser has not yet been implemented")
		}),
		FieldDeleteFieldHandler: field.DeleteFieldHandlerFunc(func(params field.DeleteFieldParams) middleware.Responder {
			return middleware.NotImplemented("operation FieldDeleteField has not yet been implemented")
		}),
		ItemDeleteItemHandler: item.DeleteItemHandlerFunc(func(params item.DeleteItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemDeleteItem has not yet been implemented")
		}),
		ItemTypeDeleteItemTypeHandler: item_type.DeleteItemTypeHandlerFunc(func(params item_type.DeleteItemTypeParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemTypeDeleteItemType has not yet been implemented")
		}),
		ProjectDeleteProjectHandler: project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation ProjectDeleteProject has not yet been implemented")
		}),
		RoleDeleteRoleHandler: role.DeleteRoleHandlerFunc(func(params role.DeleteRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation RoleDeleteRole has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserDeleteUser has not yet been implemented")
		}),
		ItemTypeDuplicateItemTypeHandler: item_type.DuplicateItemTypeHandlerFunc(func(params item_type.DuplicateItemTypeParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemTypeDuplicateItemType has not yet been implemented")
		}),
		FieldGetFieldHandler: field.GetFieldHandlerFunc(func(params field.GetFieldParams) middleware.Responder {
			return middleware.NotImplemented("operation FieldGetField has not yet been implemented")
		}),
		FieldGetFieldsHandler: field.GetFieldsHandlerFunc(func(params field.GetFieldsParams) middleware.Responder {
			return middleware.NotImplemented("operation FieldGetFields has not yet been implemented")
		}),
		ItemGetItemHandler: item.GetItemHandlerFunc(func(params item.GetItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemGetItem has not yet been implemented")
		}),
		ItemTypeGetItemTypeHandler: item_type.GetItemTypeHandlerFunc(func(params item_type.GetItemTypeParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemTypeGetItemType has not yet been implemented")
		}),
		ItemTypeGetItemTypesHandler: item_type.GetItemTypesHandlerFunc(func(params item_type.GetItemTypesParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemTypeGetItemTypes has not yet been implemented")
		}),
		ItemGetItemsHandler: item.GetItemsHandlerFunc(func(params item.GetItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemGetItems has not yet been implemented")
		}),
		ProjectGetProjectHandler: project.GetProjectHandlerFunc(func(params project.GetProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation ProjectGetProject has not yet been implemented")
		}),
		FieldGetProjectFieldsHandler: field.GetProjectFieldsHandlerFunc(func(params field.GetProjectFieldsParams) middleware.Responder {
			return middleware.NotImplemented("operation FieldGetProjectFields has not yet been implemented")
		}),
		ProjectGetProjectsHandler: project.GetProjectsHandlerFunc(func(params project.GetProjectsParams) middleware.Responder {
			return middleware.NotImplemented("operation ProjectGetProjects has not yet been implemented")
		}),
		RoleGetRoleHandler: role.GetRoleHandlerFunc(func(params role.GetRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation RoleGetRole has not yet been implemented")
		}),
		RoleGetRolesHandler: role.GetRolesHandlerFunc(func(params role.GetRolesParams) middleware.Responder {
			return middleware.NotImplemented("operation RoleGetRoles has not yet been implemented")
		}),
		UserGetUserHandler: user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserGetUser has not yet been implemented")
		}),
		UserGetUsersHandler: user.GetUsersHandlerFunc(func(params user.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation UserGetUsers has not yet been implemented")
		}),
		AccountSigninHandler: account.SigninHandlerFunc(func(params account.SigninParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountSignin has not yet been implemented")
		}),
		AccountSignupHandler: account.SignupHandlerFunc(func(params account.SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountSignup has not yet been implemented")
		}),
		FieldUpdateFieldHandler: field.UpdateFieldHandlerFunc(func(params field.UpdateFieldParams) middleware.Responder {
			return middleware.NotImplemented("operation FieldUpdateField has not yet been implemented")
		}),
		ItemUpdateItemHandler: item.UpdateItemHandlerFunc(func(params item.UpdateItemParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemUpdateItem has not yet been implemented")
		}),
		ItemTypeUpdateItemTypeHandler: item_type.UpdateItemTypeHandlerFunc(func(params item_type.UpdateItemTypeParams) middleware.Responder {
			return middleware.NotImplemented("operation ItemTypeUpdateItemType has not yet been implemented")
		}),
		ProjectUpdateProjectHandler: project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams) middleware.Responder {
			return middleware.NotImplemented("operation ProjectUpdateProject has not yet been implemented")
		}),
		RoleUpdateRoleHandler: role.UpdateRoleHandlerFunc(func(params role.UpdateRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation RoleUpdateRole has not yet been implemented")
		}),
		UserUpdateUserHandler: user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UserUpdateUser has not yet been implemented")
		}),
	}
}

/*CmsAPI the cms API */
type CmsAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// FieldCreateFieldHandler sets the operation handler for the create field operation
	FieldCreateFieldHandler field.CreateFieldHandler
	// ItemCreateItemHandler sets the operation handler for the create item operation
	ItemCreateItemHandler item.CreateItemHandler
	// ItemTypeCreateItemTypeHandler sets the operation handler for the create item type operation
	ItemTypeCreateItemTypeHandler item_type.CreateItemTypeHandler
	// ProjectCreateProjectHandler sets the operation handler for the create project operation
	ProjectCreateProjectHandler project.CreateProjectHandler
	// RoleCreateRoleHandler sets the operation handler for the create role operation
	RoleCreateRoleHandler role.CreateRoleHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// FieldDeleteFieldHandler sets the operation handler for the delete field operation
	FieldDeleteFieldHandler field.DeleteFieldHandler
	// ItemDeleteItemHandler sets the operation handler for the delete item operation
	ItemDeleteItemHandler item.DeleteItemHandler
	// ItemTypeDeleteItemTypeHandler sets the operation handler for the delete item type operation
	ItemTypeDeleteItemTypeHandler item_type.DeleteItemTypeHandler
	// ProjectDeleteProjectHandler sets the operation handler for the delete project operation
	ProjectDeleteProjectHandler project.DeleteProjectHandler
	// RoleDeleteRoleHandler sets the operation handler for the delete role operation
	RoleDeleteRoleHandler role.DeleteRoleHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// ItemTypeDuplicateItemTypeHandler sets the operation handler for the duplicate item type operation
	ItemTypeDuplicateItemTypeHandler item_type.DuplicateItemTypeHandler
	// FieldGetFieldHandler sets the operation handler for the get field operation
	FieldGetFieldHandler field.GetFieldHandler
	// FieldGetFieldsHandler sets the operation handler for the get fields operation
	FieldGetFieldsHandler field.GetFieldsHandler
	// ItemGetItemHandler sets the operation handler for the get item operation
	ItemGetItemHandler item.GetItemHandler
	// ItemTypeGetItemTypeHandler sets the operation handler for the get item type operation
	ItemTypeGetItemTypeHandler item_type.GetItemTypeHandler
	// ItemTypeGetItemTypesHandler sets the operation handler for the get item types operation
	ItemTypeGetItemTypesHandler item_type.GetItemTypesHandler
	// ItemGetItemsHandler sets the operation handler for the get items operation
	ItemGetItemsHandler item.GetItemsHandler
	// ProjectGetProjectHandler sets the operation handler for the get project operation
	ProjectGetProjectHandler project.GetProjectHandler
	// FieldGetProjectFieldsHandler sets the operation handler for the get project fields operation
	FieldGetProjectFieldsHandler field.GetProjectFieldsHandler
	// ProjectGetProjectsHandler sets the operation handler for the get projects operation
	ProjectGetProjectsHandler project.GetProjectsHandler
	// RoleGetRoleHandler sets the operation handler for the get role operation
	RoleGetRoleHandler role.GetRoleHandler
	// RoleGetRolesHandler sets the operation handler for the get roles operation
	RoleGetRolesHandler role.GetRolesHandler
	// UserGetUserHandler sets the operation handler for the get user operation
	UserGetUserHandler user.GetUserHandler
	// UserGetUsersHandler sets the operation handler for the get users operation
	UserGetUsersHandler user.GetUsersHandler
	// AccountSigninHandler sets the operation handler for the signin operation
	AccountSigninHandler account.SigninHandler
	// AccountSignupHandler sets the operation handler for the signup operation
	AccountSignupHandler account.SignupHandler
	// FieldUpdateFieldHandler sets the operation handler for the update field operation
	FieldUpdateFieldHandler field.UpdateFieldHandler
	// ItemUpdateItemHandler sets the operation handler for the update item operation
	ItemUpdateItemHandler item.UpdateItemHandler
	// ItemTypeUpdateItemTypeHandler sets the operation handler for the update item type operation
	ItemTypeUpdateItemTypeHandler item_type.UpdateItemTypeHandler
	// ProjectUpdateProjectHandler sets the operation handler for the update project operation
	ProjectUpdateProjectHandler project.UpdateProjectHandler
	// RoleUpdateRoleHandler sets the operation handler for the update role operation
	RoleUpdateRoleHandler role.UpdateRoleHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CmsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CmsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CmsAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CmsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CmsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CmsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CmsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CmsAPI
func (o *CmsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.FieldCreateFieldHandler == nil {
		unregistered = append(unregistered, "field.CreateFieldHandler")
	}

	if o.ItemCreateItemHandler == nil {
		unregistered = append(unregistered, "item.CreateItemHandler")
	}

	if o.ItemTypeCreateItemTypeHandler == nil {
		unregistered = append(unregistered, "item_type.CreateItemTypeHandler")
	}

	if o.ProjectCreateProjectHandler == nil {
		unregistered = append(unregistered, "project.CreateProjectHandler")
	}

	if o.RoleCreateRoleHandler == nil {
		unregistered = append(unregistered, "role.CreateRoleHandler")
	}

	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}

	if o.FieldDeleteFieldHandler == nil {
		unregistered = append(unregistered, "field.DeleteFieldHandler")
	}

	if o.ItemDeleteItemHandler == nil {
		unregistered = append(unregistered, "item.DeleteItemHandler")
	}

	if o.ItemTypeDeleteItemTypeHandler == nil {
		unregistered = append(unregistered, "item_type.DeleteItemTypeHandler")
	}

	if o.ProjectDeleteProjectHandler == nil {
		unregistered = append(unregistered, "project.DeleteProjectHandler")
	}

	if o.RoleDeleteRoleHandler == nil {
		unregistered = append(unregistered, "role.DeleteRoleHandler")
	}

	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}

	if o.ItemTypeDuplicateItemTypeHandler == nil {
		unregistered = append(unregistered, "item_type.DuplicateItemTypeHandler")
	}

	if o.FieldGetFieldHandler == nil {
		unregistered = append(unregistered, "field.GetFieldHandler")
	}

	if o.FieldGetFieldsHandler == nil {
		unregistered = append(unregistered, "field.GetFieldsHandler")
	}

	if o.ItemGetItemHandler == nil {
		unregistered = append(unregistered, "item.GetItemHandler")
	}

	if o.ItemTypeGetItemTypeHandler == nil {
		unregistered = append(unregistered, "item_type.GetItemTypeHandler")
	}

	if o.ItemTypeGetItemTypesHandler == nil {
		unregistered = append(unregistered, "item_type.GetItemTypesHandler")
	}

	if o.ItemGetItemsHandler == nil {
		unregistered = append(unregistered, "item.GetItemsHandler")
	}

	if o.ProjectGetProjectHandler == nil {
		unregistered = append(unregistered, "project.GetProjectHandler")
	}

	if o.FieldGetProjectFieldsHandler == nil {
		unregistered = append(unregistered, "field.GetProjectFieldsHandler")
	}

	if o.ProjectGetProjectsHandler == nil {
		unregistered = append(unregistered, "project.GetProjectsHandler")
	}

	if o.RoleGetRoleHandler == nil {
		unregistered = append(unregistered, "role.GetRoleHandler")
	}

	if o.RoleGetRolesHandler == nil {
		unregistered = append(unregistered, "role.GetRolesHandler")
	}

	if o.UserGetUserHandler == nil {
		unregistered = append(unregistered, "user.GetUserHandler")
	}

	if o.UserGetUsersHandler == nil {
		unregistered = append(unregistered, "user.GetUsersHandler")
	}

	if o.AccountSigninHandler == nil {
		unregistered = append(unregistered, "account.SigninHandler")
	}

	if o.AccountSignupHandler == nil {
		unregistered = append(unregistered, "account.SignupHandler")
	}

	if o.FieldUpdateFieldHandler == nil {
		unregistered = append(unregistered, "field.UpdateFieldHandler")
	}

	if o.ItemUpdateItemHandler == nil {
		unregistered = append(unregistered, "item.UpdateItemHandler")
	}

	if o.ItemTypeUpdateItemTypeHandler == nil {
		unregistered = append(unregistered, "item_type.UpdateItemTypeHandler")
	}

	if o.ProjectUpdateProjectHandler == nil {
		unregistered = append(unregistered, "project.UpdateProjectHandler")
	}

	if o.RoleUpdateRoleHandler == nil {
		unregistered = append(unregistered, "role.UpdateRoleHandler")
	}

	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CmsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CmsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *CmsAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *CmsAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CmsAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CmsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the cms API
func (o *CmsAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CmsAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/item-types/{itemTypeId}/fields"] = field.NewCreateField(o.context, o.FieldCreateFieldHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/items"] = item.NewCreateItem(o.context, o.ItemCreateItemHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/item-types"] = item_type.NewCreateItemType(o.context, o.ItemTypeCreateItemTypeHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects"] = project.NewCreateProject(o.context, o.ProjectCreateProjectHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/roles"] = role.NewCreateRole(o.context, o.RoleCreateRoleHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/users"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectId}/fields/{id}"] = field.NewDeleteField(o.context, o.FieldDeleteFieldHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectId}/items/{id}"] = item.NewDeleteItem(o.context, o.ItemDeleteItemHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectId}/item-types/{id}"] = item_type.NewDeleteItemType(o.context, o.ItemTypeDeleteItemTypeHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{id}"] = project.NewDeleteProject(o.context, o.ProjectDeleteProjectHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectId}/roles/{id}"] = role.NewDeleteRole(o.context, o.RoleDeleteRoleHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/projects/{projectId}/users/{id}"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/projects/{projectId}/item-types/{id}/duplicate"] = item_type.NewDuplicateItemType(o.context, o.ItemTypeDuplicateItemTypeHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/fields/{id}"] = field.NewGetField(o.context, o.FieldGetFieldHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/item-types/{itemTypeId}/fields"] = field.NewGetFields(o.context, o.FieldGetFieldsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/items/{id}"] = item.NewGetItem(o.context, o.ItemGetItemHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/item-types/{id}"] = item_type.NewGetItemType(o.context, o.ItemTypeGetItemTypeHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/item-types"] = item_type.NewGetItemTypes(o.context, o.ItemTypeGetItemTypesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/items"] = item.NewGetItems(o.context, o.ItemGetItemsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{id}"] = project.NewGetProject(o.context, o.ProjectGetProjectHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/fields"] = field.NewGetProjectFields(o.context, o.FieldGetProjectFieldsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects"] = project.NewGetProjects(o.context, o.ProjectGetProjectsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/roles/{id}"] = role.NewGetRole(o.context, o.RoleGetRoleHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/roles"] = role.NewGetRoles(o.context, o.RoleGetRolesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/users/{id}"] = user.NewGetUser(o.context, o.UserGetUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/projects/{projectId}/users"] = user.NewGetUsers(o.context, o.UserGetUsersHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/signin"] = account.NewSignin(o.context, o.AccountSigninHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/signup"] = account.NewSignup(o.context, o.AccountSignupHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectId}/fields/{id}"] = field.NewUpdateField(o.context, o.FieldUpdateFieldHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectId}/items/{id}"] = item.NewUpdateItem(o.context, o.ItemUpdateItemHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectId}/item-types/{id}"] = item_type.NewUpdateItemType(o.context, o.ItemTypeUpdateItemTypeHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{id}"] = project.NewUpdateProject(o.context, o.ProjectUpdateProjectHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectId}/roles/{id}"] = role.NewUpdateRole(o.context, o.RoleUpdateRoleHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/projects/{projectId}/users/{id}"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CmsAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CmsAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CmsAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CmsAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
