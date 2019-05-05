// Code generated by go-swagger; DO NOT EDIT.

package field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateFieldHandlerFunc turns a function with the right signature into a update field handler
type UpdateFieldHandlerFunc func(UpdateFieldParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateFieldHandlerFunc) Handle(params UpdateFieldParams) middleware.Responder {
	return fn(params)
}

// UpdateFieldHandler interface for that can handle valid update field params
type UpdateFieldHandler interface {
	Handle(UpdateFieldParams) middleware.Responder
}

// NewUpdateField creates a new http.Handler for the update field operation
func NewUpdateField(ctx *middleware.Context, handler UpdateFieldHandler) *UpdateField {
	return &UpdateField{Context: ctx, Handler: handler}
}

/*UpdateField swagger:route PUT /projects/{projectId}/fields/{id} field updateField

Update a field

*/
type UpdateField struct {
	Context *middleware.Context
	Handler UpdateFieldHandler
}

func (o *UpdateField) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateFieldParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}