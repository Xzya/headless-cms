// Code generated by go-swagger; DO NOT EDIT.

package item_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteItemTypeHandlerFunc turns a function with the right signature into a delete item type handler
type DeleteItemTypeHandlerFunc func(DeleteItemTypeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteItemTypeHandlerFunc) Handle(params DeleteItemTypeParams) middleware.Responder {
	return fn(params)
}

// DeleteItemTypeHandler interface for that can handle valid delete item type params
type DeleteItemTypeHandler interface {
	Handle(DeleteItemTypeParams) middleware.Responder
}

// NewDeleteItemType creates a new http.Handler for the delete item type operation
func NewDeleteItemType(ctx *middleware.Context, handler DeleteItemTypeHandler) *DeleteItemType {
	return &DeleteItemType{Context: ctx, Handler: handler}
}

/*DeleteItemType swagger:route DELETE /projects/{projectId}/item-types/{id} item-type deleteItemType

Delete a model

*/
type DeleteItemType struct {
	Context *middleware.Context
	Handler DeleteItemTypeHandler
}

func (o *DeleteItemType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteItemTypeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
