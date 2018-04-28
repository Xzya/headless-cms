package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

func MakeGetUsersHandler(s cms.Service) user.GetUsersHandlerFunc {
	return func(params user.GetUsersParams) middleware.Responder {
		items, err := s.GetUserList(StringToUint(params.ProjectID))
		if err != nil {
			return user.NewGetUsersDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.User, 0)
		for _, item := range items {
			data = append(data, UserToSwag(&item))
		}

		res := &apimodels.UsersResponse{
			Data: data,
		}

		return user.NewGetUsersOK().WithPayload(res)
	}
}
