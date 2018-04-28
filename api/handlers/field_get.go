package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

func MakeGetFieldHandler(s cms.Service) field.GetFieldHandlerFunc {
	return func(params field.GetFieldParams) middleware.Responder {
		req := FieldWithID(StringToUint(params.ID))
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.GetField(req); err != nil {
			return field.NewGetFieldDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.FieldResponse{
			Data: FieldToSwag(req),
		}

		return field.NewGetFieldOK().WithPayload(res)
	}
}
