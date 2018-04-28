package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

func MakeCreateFieldHandler(s cms.Service) field.CreateFieldHandlerFunc {
	return func(params field.CreateFieldParams) middleware.Responder {
		req := SwagToField(params.Body.Data)
		req.ProjectID = StringToUint(params.ProjectID)
		req.ItemTypeID = StringToUint(params.ItemTypeID)

		if err := s.CreateField(req); err != nil {
			return field.NewCreateFieldDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.FieldResponse{
			Data: FieldToSwag(req),
		}

		return field.NewCreateFieldCreated().WithPayload(res)
	}
}
