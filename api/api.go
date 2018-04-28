package api

import (
	"github.com/go-openapi/loads"
	"flag"
	"github.com/Xzya/headless-cms/gen/restapi"
	"github.com/Xzya/headless-cms/gen/restapi/operations"
	"github.com/Xzya/headless-cms"
	"github.com/Xzya/headless-cms/api/handlers"
	"github.com/go-openapi/runtime/middleware"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	gorillahandlers "github.com/gorilla/handlers"
	"os"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func Start(s cms.Service) error {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}

	api := operations.NewCmsAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	api.AccountSignupHandler = handlers.MakeCreateAccountHandler(s)
	api.AccountSigninHandler = handlers.MakeAuthenticateAccountHandler(s)

	api.ProjectCreateProjectHandler = handlers.MakeCreateProjectHandler(s)
	api.ProjectGetProjectHandler = handlers.MakeGetProjectHandler(s)
	api.ProjectUpdateProjectHandler = handlers.MakeUpdateProjectHandler(s)
	api.ProjectGetProjectsHandler = handlers.MakeGetProjectsHandler(s)
	api.ProjectDeleteProjectHandler = handlers.MakeDeleteProjectHandler(s)

	api.RoleCreateRoleHandler = handlers.MakeCreateRoleHandler(s)
	api.RoleGetRoleHandler = handlers.MakeGetRoleHandler(s)
	api.RoleUpdateRoleHandler = handlers.MakeUpdateRoleHandler(s)
	api.RoleGetRolesHandler = handlers.MakeGetRolesHandler(s)
	api.RoleDeleteRoleHandler = handlers.MakeDeleteRoleHandler(s)

	api.UserCreateUserHandler = handlers.MakeCreateUserHandler(s)
	api.UserGetUserHandler = handlers.MakeGetUserHandler(s)
	api.UserUpdateUserHandler = handlers.MakeUpdateUserHandler(s)
	api.UserGetUsersHandler = handlers.MakeGetUsersHandler(s)
	api.UserDeleteUserHandler = handlers.MakeDeleteUserHandler(s)

	api.ItemTypeCreateItemTypeHandler = handlers.MakeCreateItemTypeHandler(s)
	api.ItemTypeGetItemTypeHandler = handlers.MakeGetItemTypeHandler(s)
	api.ItemTypeUpdateItemTypeHandler = handlers.MakeUpdateItemTypeHandler(s)
	api.ItemTypeGetItemTypesHandler = handlers.MakeGetItemTypesHandler(s)
	api.ItemTypeDeleteItemTypeHandler = handlers.MakeDeleteItemTypeHandler(s)
	api.ItemTypeDuplicateItemTypeHandler = handlers.MakeDuplicateItemTypeHandler(s)

	api.FieldCreateFieldHandler = handlers.MakeCreateFieldHandler(s)
	api.FieldGetFieldHandler = handlers.MakeGetFieldHandler(s)
	api.FieldUpdateFieldHandler = handlers.MakeUpdateFieldHandler(s)
	api.FieldGetFieldsHandler = handlers.MakeGetFieldsHandler(s)
	api.FieldDeleteFieldHandler = handlers.MakeDeleteFieldHandler(s)
	api.FieldGetProjectFieldsHandler = handlers.MakeGetProjectFieldsHandler(s)

	api.ItemCreateItemHandler = handlers.MakeCreateItemHandler(s)
	api.ItemGetItemHandler = handlers.MakeGetItemHandler(s)
	api.ItemUpdateItemHandler = handlers.MakeUpdateItemHandler(s)
	api.ItemGetItemsHandler = handlers.MakeGetItemsHandler(s)
	api.ItemDeleteItemHandler = handlers.MakeDeleteItemHandler(s)

	handler := middleware.Redoc(middleware.RedocOpts{
		BasePath: "/",
		SpecURL:  swaggerSpec.SpecFilePath(),
		Path:     "/redoc",
	}, gorillahandlers.CORS(gorillahandlers.AllowedHeaders([]string{"Content-Type", "Accept", "Authorization"}),
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)(gorillahandlers.LoggingHandler(os.Stdout, api.Serve(nil))))

	server.SetHandler(handler)

	server.Port = *portFlag

	return server.Serve()
}
