// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Xzya/headless-cms/gen/models"
)

// DeleteProjectOKCode is the HTTP code returned for type DeleteProjectOK
const DeleteProjectOKCode int = 200

/*DeleteProjectOK The project details

swagger:response deleteProjectOK
*/
type DeleteProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectResponse `json:"body,omitempty"`
}

// NewDeleteProjectOK creates DeleteProjectOK with default headers values
func NewDeleteProjectOK() *DeleteProjectOK {

	return &DeleteProjectOK{}
}

// WithPayload adds the payload to the delete project o k response
func (o *DeleteProjectOK) WithPayload(payload *models.ProjectResponse) *DeleteProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project o k response
func (o *DeleteProjectOK) SetPayload(payload *models.ProjectResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteProjectDefault unexpected error

swagger:response deleteProjectDefault
*/
type DeleteProjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectDefault creates DeleteProjectDefault with default headers values
func NewDeleteProjectDefault(code int) *DeleteProjectDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteProjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete project default response
func (o *DeleteProjectDefault) WithStatusCode(code int) *DeleteProjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete project default response
func (o *DeleteProjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete project default response
func (o *DeleteProjectDefault) WithPayload(payload *models.Error) *DeleteProjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project default response
func (o *DeleteProjectDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
