package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
)

func MakeGetItemTypeHandler(s cms.Service) item_type.GetItemTypeHandlerFunc {
	return func(params item_type.GetItemTypeParams) middleware.Responder {
		req := ItemTypeWithID(StringToUint(params.ID))
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.GetItemType(req); err != nil {
			return item_type.NewGetItemTypeDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemTypeResponse{
			Data:     ItemTypeToSwag(req),
			Included: ItemTypeIncluded(req),
		}

		return item_type.NewGetItemTypeOK().WithPayload(res)
	}
}
