package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
)

func MakeUpdateItemHandler(s cms.Service) item.UpdateItemHandlerFunc {
	return func(params item.UpdateItemParams) middleware.Responder {
		req := SwagToItem(params.Body.Data)
		req.ID = StringToUint(params.ID)
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.UpdateItem(req); err != nil {
			return item.NewUpdateItemDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemResponse{
			Data: ItemToSwag(req),
		}

		return item.NewUpdateItemOK().WithPayload(res)
	}
}
