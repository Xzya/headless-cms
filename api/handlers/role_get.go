package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
)

func MakeGetRoleHandler(s cms.Service) role.GetRoleHandlerFunc {
	return func(params role.GetRoleParams) middleware.Responder {
		req := RoleWithID(StringToUint(params.ID))
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.GetRole(req); err != nil {
			return role.NewGetRoleDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.RoleResponse{
			Data: RoleToSwag(req),
		}

		return role.NewGetRoleOK().WithPayload(res)
	}
}
