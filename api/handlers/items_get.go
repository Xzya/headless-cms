package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item"
)

func MakeGetItemsHandler(s cms.Service) item.GetItemsHandlerFunc {
	return func(params item.GetItemsParams) middleware.Responder {
		count, items, err := s.GetItemList(StringToUint(params.ProjectID), SwagToItemFilter(params))
		if err != nil {
			return item.NewGetItemsDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.Item, 0)
		for _, item := range items {
			data = append(data, ItemToSwag(&item))
		}

		res := &apimodels.ItemsResponse{
			Data: data,
			Meta: &apimodels.ItemsResponseMeta{
				TotalCount: int32(count),
			},
		}

		return item.NewGetItemsOK().WithPayload(res)
	}
}
