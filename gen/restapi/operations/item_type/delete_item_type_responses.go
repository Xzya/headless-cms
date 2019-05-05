// Code generated by go-swagger; DO NOT EDIT.

package item_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Xzya/headless-cms/gen/models"
)

// DeleteItemTypeOKCode is the HTTP code returned for type DeleteItemTypeOK
const DeleteItemTypeOKCode int = 200

/*DeleteItemTypeOK The model details

swagger:response deleteItemTypeOK
*/
type DeleteItemTypeOK struct {

	/*
	  In: Body
	*/
	Payload *models.ItemTypeResponse `json:"body,omitempty"`
}

// NewDeleteItemTypeOK creates DeleteItemTypeOK with default headers values
func NewDeleteItemTypeOK() *DeleteItemTypeOK {

	return &DeleteItemTypeOK{}
}

// WithPayload adds the payload to the delete item type o k response
func (o *DeleteItemTypeOK) WithPayload(payload *models.ItemTypeResponse) *DeleteItemTypeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete item type o k response
func (o *DeleteItemTypeOK) SetPayload(payload *models.ItemTypeResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteItemTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteItemTypeDefault unexpected error

swagger:response deleteItemTypeDefault
*/
type DeleteItemTypeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteItemTypeDefault creates DeleteItemTypeDefault with default headers values
func NewDeleteItemTypeDefault(code int) *DeleteItemTypeDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteItemTypeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete item type default response
func (o *DeleteItemTypeDefault) WithStatusCode(code int) *DeleteItemTypeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete item type default response
func (o *DeleteItemTypeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete item type default response
func (o *DeleteItemTypeDefault) WithPayload(payload *models.Error) *DeleteItemTypeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete item type default response
func (o *DeleteItemTypeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteItemTypeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}