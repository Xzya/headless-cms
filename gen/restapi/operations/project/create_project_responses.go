// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Xzya/headless-cms/gen/models"
)

// CreateProjectCreatedCode is the HTTP code returned for type CreateProjectCreated
const CreateProjectCreatedCode int = 201

/*CreateProjectCreated Created

swagger:response createProjectCreated
*/
type CreateProjectCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectResponse `json:"body,omitempty"`
}

// NewCreateProjectCreated creates CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {

	return &CreateProjectCreated{}
}

// WithPayload adds the payload to the create project created response
func (o *CreateProjectCreated) WithPayload(payload *models.ProjectResponse) *CreateProjectCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create project created response
func (o *CreateProjectCreated) SetPayload(payload *models.ProjectResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProjectCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateProjectDefault unexpected error

swagger:response createProjectDefault
*/
type CreateProjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateProjectDefault creates CreateProjectDefault with default headers values
func NewCreateProjectDefault(code int) *CreateProjectDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateProjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create project default response
func (o *CreateProjectDefault) WithStatusCode(code int) *CreateProjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create project default response
func (o *CreateProjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create project default response
func (o *CreateProjectDefault) WithPayload(payload *models.Error) *CreateProjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create project default response
func (o *CreateProjectDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateProjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}