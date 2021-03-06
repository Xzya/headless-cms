// Code generated by go-swagger; DO NOT EDIT.

package item_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDuplicateItemTypeParams creates a new DuplicateItemTypeParams object
// no default values defined in spec.
func NewDuplicateItemTypeParams() DuplicateItemTypeParams {

	return DuplicateItemTypeParams{}
}

// DuplicateItemTypeParams contains all the bound params for the duplicate item type operation
// typically these are obtained from a http.Request
//
// swagger:parameters duplicateItemType
type DuplicateItemTypeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The model id
	  Required: true
	  In: path
	*/
	ID string
	/*The project id
	  Required: true
	  In: path
	*/
	ProjectID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDuplicateItemTypeParams() beforehand.
func (o *DuplicateItemTypeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rProjectID, rhkProjectID, _ := route.Params.GetOK("projectId")
	if err := o.bindProjectID(rProjectID, rhkProjectID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DuplicateItemTypeParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

func (o *DuplicateItemTypeParams) bindProjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectID = raw

	return nil
}
