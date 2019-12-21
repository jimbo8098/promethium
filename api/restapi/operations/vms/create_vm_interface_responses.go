// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// CreateVMInterfaceOKCode is the HTTP code returned for type CreateVMInterfaceOK
const CreateVMInterfaceOKCode int = 200

/*CreateVMInterfaceOK successful operation

swagger:response createVmInterfaceOK
*/
type CreateVMInterfaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.MetaDataNetworkInterfaceConfig `json:"body,omitempty"`
}

// NewCreateVMInterfaceOK creates CreateVMInterfaceOK with default headers values
func NewCreateVMInterfaceOK() *CreateVMInterfaceOK {

	return &CreateVMInterfaceOK{}
}

// WithPayload adds the payload to the create Vm interface o k response
func (o *CreateVMInterfaceOK) WithPayload(payload *models.MetaDataNetworkInterfaceConfig) *CreateVMInterfaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Vm interface o k response
func (o *CreateVMInterfaceOK) SetPayload(payload *models.MetaDataNetworkInterfaceConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVMInterfaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVMInterfaceBadRequestCode is the HTTP code returned for type CreateVMInterfaceBadRequest
const CreateVMInterfaceBadRequestCode int = 400

/*CreateVMInterfaceBadRequest Invalid ID supplied

swagger:response createVmInterfaceBadRequest
*/
type CreateVMInterfaceBadRequest struct {
}

// NewCreateVMInterfaceBadRequest creates CreateVMInterfaceBadRequest with default headers values
func NewCreateVMInterfaceBadRequest() *CreateVMInterfaceBadRequest {

	return &CreateVMInterfaceBadRequest{}
}

// WriteResponse to the client
func (o *CreateVMInterfaceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateVMInterfaceNotFoundCode is the HTTP code returned for type CreateVMInterfaceNotFound
const CreateVMInterfaceNotFoundCode int = 404

/*CreateVMInterfaceNotFound VM not found

swagger:response createVmInterfaceNotFound
*/
type CreateVMInterfaceNotFound struct {
}

// NewCreateVMInterfaceNotFound creates CreateVMInterfaceNotFound with default headers values
func NewCreateVMInterfaceNotFound() *CreateVMInterfaceNotFound {

	return &CreateVMInterfaceNotFound{}
}

// WriteResponse to the client
func (o *CreateVMInterfaceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*CreateVMInterfaceDefault unexpected error

swagger:response createVmInterfaceDefault
*/
type CreateVMInterfaceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateVMInterfaceDefault creates CreateVMInterfaceDefault with default headers values
func NewCreateVMInterfaceDefault(code int) *CreateVMInterfaceDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateVMInterfaceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create VM interface default response
func (o *CreateVMInterfaceDefault) WithStatusCode(code int) *CreateVMInterfaceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create VM interface default response
func (o *CreateVMInterfaceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create VM interface default response
func (o *CreateVMInterfaceDefault) WithPayload(payload *models.Error) *CreateVMInterfaceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create VM interface default response
func (o *CreateVMInterfaceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVMInterfaceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
