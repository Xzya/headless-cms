package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
)

func MakeGetProjectHandler(s cms.Service) project.GetProjectHandlerFunc {
	return func(params project.GetProjectParams) middleware.Responder {
		req := ProjectWithID(StringToUint(params.ID))
		// TODO: - return only the user's projects

		if err := s.GetProject(req); err != nil {
			return project.NewGetProjectDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ProjectResponse{
			Data: ProjectToSwag(req),
		}

		return project.NewGetProjectOK().WithPayload(res)
	}
}
