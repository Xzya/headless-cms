package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
)

func MakeDeleteProjectHandler(s cms.Service) project.DeleteProjectHandlerFunc {
	return func(params project.DeleteProjectParams) middleware.Responder {
		req := ProjectWithID(StringToUint(params.ID))

		if err := s.DeleteProject(req); err != nil {
			return project.NewDeleteProjectDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ProjectResponse{
			Data: ProjectToSwag(req),
		}

		return project.NewDeleteProjectOK().WithPayload(res)
	}
}
