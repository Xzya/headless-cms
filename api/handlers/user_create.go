package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

func MakeCreateUserHandler(s cms.Service) user.CreateUserHandlerFunc {
	return func(params user.CreateUserParams) middleware.Responder {
		req := SwagToUser(params.Body.Data)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.CreateUser(req); err != nil {
			return user.NewCreateUserDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.UserResponse{
			Data: UserToSwag(req),
		}

		return user.NewCreateUserCreated().WithPayload(res)
	}
}
