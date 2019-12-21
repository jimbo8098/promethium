// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// CreateImageOKCode is the HTTP code returned for type CreateImageOK
const CreateImageOKCode int = 200

/*CreateImageOK successful operation

swagger:response createImageOK
*/
type CreateImageOK struct {

	/*
	  In: Body
	*/
	Payload *models.VM `json:"body,omitempty"`
}

// NewCreateImageOK creates CreateImageOK with default headers values
func NewCreateImageOK() *CreateImageOK {

	return &CreateImageOK{}
}

// WithPayload adds the payload to the create image o k response
func (o *CreateImageOK) WithPayload(payload *models.VM) *CreateImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create image o k response
func (o *CreateImageOK) SetPayload(payload *models.VM) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateImageDefault unexpected error

swagger:response createImageDefault
*/
type CreateImageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateImageDefault creates CreateImageDefault with default headers values
func NewCreateImageDefault(code int) *CreateImageDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateImageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create image default response
func (o *CreateImageDefault) WithStatusCode(code int) *CreateImageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create image default response
func (o *CreateImageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create image default response
func (o *CreateImageDefault) WithPayload(payload *models.Error) *CreateImageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create image default response
func (o *CreateImageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateImageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
