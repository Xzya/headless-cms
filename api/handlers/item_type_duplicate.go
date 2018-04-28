package handlers

import (
	"github.com/Xzya/headless-cms"
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Xzya/headless-cms/gen/restapi/operations/item_type"
)

func MakeDuplicateItemTypeHandler(s cms.Service) item_type.DuplicateItemTypeHandlerFunc {
	return func(params item_type.DuplicateItemTypeParams) middleware.Responder {
		req := ItemTypeWithID(StringToUint(params.ID))
		req.ProjectID = StringToUint(params.ProjectID)

		if err := s.DuplicateItemType(req); err != nil {
			return item_type.NewDuplicateItemTypeDefault(400).WithPayload(APIError(err))
		}

		res := &apimodels.ItemTypeResponse{
			Data: ItemTypeToSwag(req),
		}

		return item_type.NewDuplicateItemTypeOK().WithPayload(res)
	}
}
