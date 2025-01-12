// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// StartVMOKCode is the HTTP code returned for type StartVMOK
const StartVMOKCode int = 200

/*StartVMOK successful operation

swagger:response startVmOK
*/
type StartVMOK struct {

	/*
	  In: Body
	*/
	Payload *models.VM `json:"body,omitempty"`
}

// NewStartVMOK creates StartVMOK with default headers values
func NewStartVMOK() *StartVMOK {

	return &StartVMOK{}
}

// WithPayload adds the payload to the start Vm o k response
func (o *StartVMOK) WithPayload(payload *models.VM) *StartVMOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start Vm o k response
func (o *StartVMOK) SetPayload(payload *models.VM) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartVMOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StartVMBadRequestCode is the HTTP code returned for type StartVMBadRequest
const StartVMBadRequestCode int = 400

/*StartVMBadRequest Invalid ID supplied

swagger:response startVmBadRequest
*/
type StartVMBadRequest struct {
}

// NewStartVMBadRequest creates StartVMBadRequest with default headers values
func NewStartVMBadRequest() *StartVMBadRequest {

	return &StartVMBadRequest{}
}

// WriteResponse to the client
func (o *StartVMBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// StartVMNotFoundCode is the HTTP code returned for type StartVMNotFound
const StartVMNotFoundCode int = 404

/*StartVMNotFound VM not found

swagger:response startVmNotFound
*/
type StartVMNotFound struct {
}

// NewStartVMNotFound creates StartVMNotFound with default headers values
func NewStartVMNotFound() *StartVMNotFound {

	return &StartVMNotFound{}
}

// WriteResponse to the client
func (o *StartVMNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
