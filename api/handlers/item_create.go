package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
)

func MakeCreateItemHandler(s cms.Service) item.CreateItemHandlerFunc {
	return func(params item.CreateItemParams) middleware.Responder {
		req := SwagToItem(params.Body.Data)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.CreateItem(req); err != nil {
			return item.NewCreateItemDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemResponse{
			Data: ItemToSwag(req),
		}

		return item.NewCreateItemCreated().WithPayload(res)
	}
}
