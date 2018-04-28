package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/project"
)

func MakeGetProjectsHandler(s cms.Service) project.GetProjectsHandlerFunc {
	return func(params project.GetProjectsParams) middleware.Responder {
		items, err := s.GetProjectList()
		if err != nil {
			return project.NewGetProjectsDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.Project, 0)
		for _, item := range items {
			data = append(data, ProjectToSwag(&item))
		}

		res := &apimodels.ProjectsResponse{
			Data: data,
		}

		return project.NewGetProjectsOK().WithPayload(res)
	}
}
