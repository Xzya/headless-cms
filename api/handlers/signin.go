package handlers

import (
	"github.com/Xzya/headless-cms"
	"github.com/Xzya/headless-cms/models"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/account"
)

func MakeAuthenticateAccountHandler(s cms.Service) account.SigninHandlerFunc {
	return func(params account.SigninParams) middleware.Responder {
		acc := &models.Account{
			Email:    params.Body.Data.Attributes.Email.String(),
			Password: params.Body.Data.Attributes.Password.String(),
		}

		token, err := s.AuthenticateAccount(acc)
		if err != nil {
			return account.NewSigninDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.SigninOKBody{
			Data: &apimodels.Session{
				ID:   token,
				Type: apimodels.SessionTypeSession,
				Relationships: &apimodels.SessionRelationships{
					User: &apimodels.SessionRelationshipsUser{
						Data: AccountToRel(acc),
					},
				},
			},
		}

		return account.NewSigninOK().WithPayload(res)
	}
}
