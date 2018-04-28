package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

func MakeUpdateUserHandler(s cms.Service) user.UpdateUserHandlerFunc {
	return func(params user.UpdateUserParams) middleware.Responder {
		req := SwagToUser(params.Body.Data)
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.UpdateUser(req); err != nil {
			return user.NewUpdateUserDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.UserResponse{
			Data: UserToSwag(req),
		}

		return user.NewUpdateUserOK().WithPayload(res)
	}
}
