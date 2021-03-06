// Code generated by go-swagger; DO NOT EDIT.

package item_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateItemTypeHandlerFunc turns a function with the right signature into a update item type handler
type UpdateItemTypeHandlerFunc func(UpdateItemTypeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateItemTypeHandlerFunc) Handle(params UpdateItemTypeParams) middleware.Responder {
	return fn(params)
}

// UpdateItemTypeHandler interface for that can handle valid update item type params
type UpdateItemTypeHandler interface {
	Handle(UpdateItemTypeParams) middleware.Responder
}

// NewUpdateItemType creates a new http.Handler for the update item type operation
func NewUpdateItemType(ctx *middleware.Context, handler UpdateItemTypeHandler) *UpdateItemType {
	return &UpdateItemType{Context: ctx, Handler: handler}
}

/*UpdateItemType swagger:route PUT /projects/{projectId}/item-types/{id} item-type updateItemType

Update a model

*/
type UpdateItemType struct {
	Context *middleware.Context
	Handler UpdateItemTypeHandler
}

func (o *UpdateItemType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateItemTypeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
