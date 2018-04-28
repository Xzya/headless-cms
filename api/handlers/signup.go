package handlers

import (
	"github.com/Xzya/headless-cms"
	"github.com/Xzya/headless-cms/models"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/account"
)

func MakeCreateAccountHandler(s cms.Service) account.SignupHandlerFunc {
	return func(params account.SignupParams) middleware.Responder {
		acc := &models.Account{
			Email:    params.Body.Data.Attributes.Email.String(),
			Password: params.Body.Data.Attributes.Password.String(),
		}

		token, err := s.CreateAccount(acc)
		if err != nil {
			return account.NewSignupDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.SignupCreatedBody{
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

		return account.NewSignupCreated().WithPayload(res)
	}
}
