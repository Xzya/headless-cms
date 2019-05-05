// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/Xzya/headless-cms/gen/restapi/operations"
	"github.com/Xzya/headless-cms/gen/restapi/operations/account"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

//go:generate swagger generate server --target ../gen --name cms --spec ../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.CmsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CmsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.FieldCreateFieldHandler = field.CreateFieldHandlerFunc(func(params field.CreateFieldParams) middleware.Responder {
		return middleware.NotImplemented("operation field.CreateField has not yet been implemented")
	})
	api.ItemCreateItemHandler = item.CreateItemHandlerFunc(func(params item.CreateItemParams) middleware.Responder {
		return middleware.NotImplemented("operation item.CreateItem has not yet been implemented")
	})
	api.ItemTypeCreateItemTypeHandler = item_type.CreateItemTypeHandlerFunc(func(params item_type.CreateItemTypeParams) middleware.Responder {
		return middleware.NotImplemented("operation item_type.CreateItemType has not yet been implemented")
	})
	api.ProjectCreateProjectHandler = project.CreateProjectHandlerFunc(func(params project.CreateProjectParams) middleware.Responder {
		return middleware.NotImplemented("operation project.CreateProject has not yet been implemented")
	})
	api.RoleCreateRoleHandler = role.CreateRoleHandlerFunc(func(params role.CreateRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation role.CreateRole has not yet been implemented")
	})
	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
	})
	api.FieldDeleteFieldHandler = field.DeleteFieldHandlerFunc(func(params field.DeleteFieldParams) middleware.Responder {
		return middleware.NotImplemented("operation field.DeleteField has not yet been implemented")
	})
	api.ItemDeleteItemHandler = item.DeleteItemHandlerFunc(func(params item.DeleteItemParams) middleware.Responder {
		return middleware.NotImplemented("operation item.DeleteItem has not yet been implemented")
	})
	api.ItemTypeDeleteItemTypeHandler = item_type.DeleteItemTypeHandlerFunc(func(params item_type.DeleteItemTypeParams) middleware.Responder {
		return middleware.NotImplemented("operation item_type.DeleteItemType has not yet been implemented")
	})
	api.ProjectDeleteProjectHandler = project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams) middleware.Responder {
		return middleware.NotImplemented("operation project.DeleteProject has not yet been implemented")
	})
	api.RoleDeleteRoleHandler = role.DeleteRoleHandlerFunc(func(params role.DeleteRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation role.DeleteRole has not yet been implemented")
	})
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
	})
	api.ItemTypeDuplicateItemTypeHandler = item_type.DuplicateItemTypeHandlerFunc(func(params item_type.DuplicateItemTypeParams) middleware.Responder {
		return middleware.NotImplemented("operation item_type.DuplicateItemType has not yet been implemented")
	})
	api.FieldGetFieldHandler = field.GetFieldHandlerFunc(func(params field.GetFieldParams) middleware.Responder {
		return middleware.NotImplemented("operation field.GetField has not yet been implemented")
	})
	api.FieldGetFieldsHandler = field.GetFieldsHandlerFunc(func(params field.GetFieldsParams) middleware.Responder {
		return middleware.NotImplemented("operation field.GetFields has not yet been implemented")
	})
	api.ItemGetItemHandler = item.GetItemHandlerFunc(func(params item.GetItemParams) middleware.Responder {
		return middleware.NotImplemented("operation item.GetItem has not yet been implemented")
	})
	api.ItemTypeGetItemTypeHandler = item_type.GetItemTypeHandlerFunc(func(params item_type.GetItemTypeParams) middleware.Responder {
		return middleware.NotImplemented("operation item_type.GetItemType has not yet been implemented")
	})
	api.ItemTypeGetItemTypesHandler = item_type.GetItemTypesHandlerFunc(func(params item_type.GetItemTypesParams) middleware.Responder {
		return middleware.NotImplemented("operation item_type.GetItemTypes has not yet been implemented")
	})
	api.ItemGetItemsHandler = item.GetItemsHandlerFunc(func(params item.GetItemsParams) middleware.Responder {
		return middleware.NotImplemented("operation item.GetItems has not yet been implemented")
	})
	api.ProjectGetProjectHandler = project.GetProjectHandlerFunc(func(params project.GetProjectParams) middleware.Responder {
		return middleware.NotImplemented("operation project.GetProject has not yet been implemented")
	})
	api.FieldGetProjectFieldsHandler = field.GetProjectFieldsHandlerFunc(func(params field.GetProjectFieldsParams) middleware.Responder {
		return middleware.NotImplemented("operation field.GetProjectFields has not yet been implemented")
	})
	api.ProjectGetProjectsHandler = project.GetProjectsHandlerFunc(func(params project.GetProjectsParams) middleware.Responder {
		return middleware.NotImplemented("operation project.GetProjects has not yet been implemented")
	})
	api.RoleGetRoleHandler = role.GetRoleHandlerFunc(func(params role.GetRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation role.GetRole has not yet been implemented")
	})
	api.RoleGetRolesHandler = role.GetRolesHandlerFunc(func(params role.GetRolesParams) middleware.Responder {
		return middleware.NotImplemented("operation role.GetRoles has not yet been implemented")
	})
	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
	})
	api.UserGetUsersHandler = user.GetUsersHandlerFunc(func(params user.GetUsersParams) middleware.Responder {
		return middleware.NotImplemented("operation user.GetUsers has not yet been implemented")
	})
	api.AccountSigninHandler = account.SigninHandlerFunc(func(params account.SigninParams) middleware.Responder {
		return middleware.NotImplemented("operation account.Signin has not yet been implemented")
	})
	api.AccountSignupHandler = account.SignupHandlerFunc(func(params account.SignupParams) middleware.Responder {
		return middleware.NotImplemented("operation account.Signup has not yet been implemented")
	})
	api.FieldUpdateFieldHandler = field.UpdateFieldHandlerFunc(func(params field.UpdateFieldParams) middleware.Responder {
		return middleware.NotImplemented("operation field.UpdateField has not yet been implemented")
	})
	api.ItemUpdateItemHandler = item.UpdateItemHandlerFunc(func(params item.UpdateItemParams) middleware.Responder {
		return middleware.NotImplemented("operation item.UpdateItem has not yet been implemented")
	})
	api.ItemTypeUpdateItemTypeHandler = item_type.UpdateItemTypeHandlerFunc(func(params item_type.UpdateItemTypeParams) middleware.Responder {
		return middleware.NotImplemented("operation item_type.UpdateItemType has not yet been implemented")
	})
	api.ProjectUpdateProjectHandler = project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams) middleware.Responder {
		return middleware.NotImplemented("operation project.UpdateProject has not yet been implemented")
	})
	api.RoleUpdateRoleHandler = role.UpdateRoleHandlerFunc(func(params role.UpdateRoleParams) middleware.Responder {
		return middleware.NotImplemented("operation role.UpdateRole has not yet been implemented")
	})
	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
		return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
