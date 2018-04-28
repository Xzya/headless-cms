package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/user"
)

func MakeDeleteUserHandler(s cms.Service) user.DeleteUserHandlerFunc {
	return func(params user.DeleteUserParams) middleware.Responder {
		req := UserWithID(StringToUint(params.ID))
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.DeleteUser(req); err != nil {
			return user.NewDeleteUserDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.UserResponse{
			Data: UserToSwag(req),
		}

		return user.NewDeleteUserOK().WithPayload(res)
	}
}
