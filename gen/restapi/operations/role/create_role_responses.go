// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Xzya/headless-cms/gen/models"
)

// CreateRoleCreatedCode is the HTTP code returned for type CreateRoleCreated
const CreateRoleCreatedCode int = 201

/*CreateRoleCreated Created

swagger:response createRoleCreated
*/
type CreateRoleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RoleResponse `json:"body,omitempty"`
}

// NewCreateRoleCreated creates CreateRoleCreated with default headers values
func NewCreateRoleCreated() *CreateRoleCreated {

	return &CreateRoleCreated{}
}

// WithPayload adds the payload to the create role created response
func (o *CreateRoleCreated) WithPayload(payload *models.RoleResponse) *CreateRoleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create role created response
func (o *CreateRoleCreated) SetPayload(payload *models.RoleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRoleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateRoleDefault unexpected error

swagger:response createRoleDefault
*/
type CreateRoleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateRoleDefault creates CreateRoleDefault with default headers values
func NewCreateRoleDefault(code int) *CreateRoleDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateRoleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create role default response
func (o *CreateRoleDefault) WithStatusCode(code int) *CreateRoleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create role default response
func (o *CreateRoleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create role default response
func (o *CreateRoleDefault) WithPayload(payload *models.Error) *CreateRoleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create role default response
func (o *CreateRoleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRoleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
