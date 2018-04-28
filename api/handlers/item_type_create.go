package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
)

func MakeCreateItemTypeHandler(s cms.Service) item_type.CreateItemTypeHandlerFunc {
	return func(params item_type.CreateItemTypeParams) middleware.Responder {
		req := SwagToItemType(params.Body.Data)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.CreateItemType(req); err != nil {
			return item_type.NewCreateItemTypeDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemTypeResponse{
			Data: ItemTypeToSwag(req),
		}

		return item_type.NewCreateItemTypeCreated().WithPayload(res)
	}
}
