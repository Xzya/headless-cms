package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

func MakeGetProjectFieldsHandler(s cms.Service) field.GetProjectFieldsHandlerFunc {
	return func(params field.GetProjectFieldsParams) middleware.Responder {
		items, err := s.GetProjectFieldList(StringToUint(params.ProjectID), SwagToFieldFilter(params))
		if err != nil {
			return field.NewGetProjectFieldsDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.Field, 0)
		for _, item := range items {
			data = append(data, FieldToSwag(&item))
		}

		res := &apimodels.FieldsResponse{
			Data: data,
		}

		return field.NewGetProjectFieldsOK().WithPayload(res)
	}
}
