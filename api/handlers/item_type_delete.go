package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
)

func MakeDeleteItemTypeHandler(s cms.Service) item_type.DeleteItemTypeHandlerFunc {
	return func(params item_type.DeleteItemTypeParams) middleware.Responder {
		req := ItemTypeWithID(StringToUint(params.ID))
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.DeleteItemType(req); err != nil {
			return item_type.NewDeleteItemTypeDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemTypeResponse{
			Data: ItemTypeToSwag(req),
		}

		return item_type.NewDeleteItemTypeOK().WithPayload(res)
	}
}
