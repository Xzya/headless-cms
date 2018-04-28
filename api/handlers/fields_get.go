package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

func MakeGetFieldsHandler(s cms.Service) field.GetFieldsHandlerFunc {
	return func(params field.GetFieldsParams) middleware.Responder {
		items, err := s.GetFieldList(StringToUint(params.ItemTypeID), StringToUint(params.ProjectID))
		if err != nil {
			return field.NewGetFieldsDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.Field, 0)
		for _, item := range items {
			data = append(data, FieldToSwag(&item))
		}

		res := &apimodels.FieldsResponse{
			Data: data,
		}

		return field.NewGetFieldsOK().WithPayload(res)
	}
}
