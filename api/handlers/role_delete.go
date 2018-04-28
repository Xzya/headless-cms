package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
)

func MakeDeleteRoleHandler(s cms.Service) role.DeleteRoleHandlerFunc {
	return func(params role.DeleteRoleParams) middleware.Responder {
		req := RoleWithID(StringToUint(params.ID))
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.DeleteRole(req); err != nil {
			return role.NewDeleteRoleDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.RoleResponse{
			Data: RoleToSwag(req),
		}

		return role.NewDeleteRoleOK().WithPayload(res)
	}
}
