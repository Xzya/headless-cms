package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/role"
)

func MakeGetRolesHandler(s cms.Service) role.GetRolesHandlerFunc {
	return func(params role.GetRolesParams) middleware.Responder {
		items, err := s.GetRoleList(StringToUint(params.ProjectID))
		if err != nil {
			return role.NewGetRolesDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.Role, 0)
		for _, item := range items {
			data = append(data, RoleToSwag(&item))
		}

		res := &apimodels.RolesResponse{
			Data: data,
		}

		return role.NewGetRolesOK().WithPayload(res)
	}
}
