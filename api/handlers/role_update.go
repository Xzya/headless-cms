package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
)

func MakeUpdateRoleHandler(s cms.Service) role.UpdateRoleHandlerFunc {
	return func(params role.UpdateRoleParams) middleware.Responder {
		req := SwagToRole(params.Body.Data)
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.UpdateRole(req); err != nil {
			return role.NewUpdateRoleDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.RoleResponse{
			Data: RoleToSwag(req),
		}

		return role.NewUpdateRoleOK().WithPayload(res)
	}
}
