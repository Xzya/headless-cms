package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
)

func MakeCreateRoleHandler(s cms.Service) role.CreateRoleHandlerFunc {
	return func(params role.CreateRoleParams) middleware.Responder {
		req := SwagToRole(params.Body.Data)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.CreateRole(req); err != nil {
			return role.NewCreateRoleDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.RoleResponse{
			Data: RoleToSwag(req),
		}

		return role.NewCreateRoleCreated().WithPayload(res)
	}
}
