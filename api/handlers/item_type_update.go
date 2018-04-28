package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
)

func MakeUpdateItemTypeHandler(s cms.Service) item_type.UpdateItemTypeHandlerFunc {
	return func(params item_type.UpdateItemTypeParams) middleware.Responder {
		req := SwagToItemType(params.Body.Data)
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.UpdateItemType(req); err != nil {
			return item_type.NewUpdateItemTypeDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemTypeResponse{
			Data: ItemTypeToSwag(req),
		}

		return item_type.NewUpdateItemTypeOK().WithPayload(res)
	}
}
