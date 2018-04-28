package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
)

func MakeGetItemHandler(s cms.Service) item.GetItemHandlerFunc {
	return func(params item.GetItemParams) middleware.Responder {
		req := ItemWithID(StringToUint(params.ID))
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.GetItem(req); err != nil {
			return item.NewGetItemDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemResponse{
			Data: ItemToSwag(req),
		}

		return item.NewGetItemOK().WithPayload(res)
	}
}
