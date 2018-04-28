package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
)

func MakeGetItemTypesHandler(s cms.Service) item_type.GetItemTypesHandlerFunc {
	return func(params item_type.GetItemTypesParams) middleware.Responder {
		items, err := s.GetItemTypeList(StringToUint(params.ProjectID))
		if err != nil {
			return item_type.NewGetItemTypesDefault(400).WithPayload(APIError(err))
		}

		data := make([]*apimodels.ItemType, 0)
		for _, item := range items {
			data = append(data, ItemTypeToSwag(&item))
		}

		res := &apimodels.ItemTypesResponse{
			Data: data,
		}

		return item_type.NewGetItemTypesOK().WithPayload(res)
	}
}
