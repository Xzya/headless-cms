package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
)

func MakeUpdateProjectHandler(s cms.Service) project.UpdateProjectHandlerFunc {
	return func(params project.UpdateProjectParams) middleware.Responder {
		req := SwagToProject(params.Body.Data)
		req.ID = StringToUint(params.ID)

		if err := s.UpdateProject(req); err != nil {
			return project.NewUpdateProjectDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ProjectResponse{
			Data: ProjectToSwag(req),
		}

		return project.NewUpdateProjectOK().WithPayload(res)
	}
}
