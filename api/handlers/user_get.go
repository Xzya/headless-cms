package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

func MakeGetUserHandler(s cms.Service) user.GetUserHandlerFunc {
	return func(params user.GetUserParams) middleware.Responder {
		req := UserWithID(StringToUint(params.ID))
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.GetUser(req); err != nil {
			return user.NewGetUserDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.UserResponse{
			Data: UserToSwag(req),
		}

		return user.NewGetUserOK().WithPayload(res)
	}
}
