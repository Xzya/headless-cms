// Code generated by go-swagger; DO NOT EDIT.

package field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteFieldHandlerFunc turns a function with the right signature into a delete field handler
type DeleteFieldHandlerFunc func(DeleteFieldParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFieldHandlerFunc) Handle(params DeleteFieldParams) middleware.Responder {
	return fn(params)
}

// DeleteFieldHandler interface for that can handle valid delete field params
type DeleteFieldHandler interface {
	Handle(DeleteFieldParams) middleware.Responder
}

// NewDeleteField creates a new http.Handler for the delete field operation
func NewDeleteField(ctx *middleware.Context, handler DeleteFieldHandler) *DeleteField {
	return &DeleteField{Context: ctx, Handler: handler}
}

/*DeleteField swagger:route DELETE /projects/{projectId}/fields/{id} field deleteField

Delete a field

*/
type DeleteField struct {
	Context *middleware.Context
	Handler DeleteFieldHandler
}

func (o *DeleteField) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFieldParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
