package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
)

func MakeDeleteItemHandler(s cms.Service) item.DeleteItemHandlerFunc {
	return func(params item.DeleteItemParams) middleware.Responder {
		req := ItemWithID(StringToUint(params.ID))
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.DeleteItem(req); err != nil {
			return item.NewDeleteItemDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemResponse{
			Data: ItemToSwag(req),
		}

		return item.NewDeleteItemOK().WithPayload(res)
	}
}
