package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

func MakeUpdateFieldHandler(s cms.Service) field.UpdateFieldHandlerFunc {
	return func(params field.UpdateFieldParams) middleware.Responder {
		req := SwagToField(params.Body.Data)
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.UpdateField(req); err != nil {
			return field.NewUpdateFieldDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.FieldResponse{
			Data: FieldToSwag(req),
		}

		return field.NewUpdateFieldOK().WithPayload(res)
	}
}
