package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/field"
)

func MakeDeleteFieldHandler(s cms.Service) field.DeleteFieldHandlerFunc {
	return func(params field.DeleteFieldParams) middleware.Responder {
		req := FieldWithID(StringToUint(params.ID))
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.DeleteField(req); err != nil {
			return field.NewDeleteFieldDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.FieldResponse{
			Data: FieldToSwag(req),
		}

		return field.NewDeleteFieldOK().WithPayload(res)
	}
}
