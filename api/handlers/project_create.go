package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
)

func MakeCreateProjectHandler(s cms.Service) project.CreateProjectHandlerFunc {
	return func(params project.CreateProjectParams) middleware.Responder {
		req := SwagToProject(params.Body.Data)
		// TODO: - get account id from token
		req.AccountID = 1

		if err := s.CreateProject(req); err != nil {
			return project.NewCreateProjectDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ProjectResponse{
			Data: ProjectToSwag(req),
		}

		return project.NewCreateProjectCreated().WithPayload(res)
	}
}
