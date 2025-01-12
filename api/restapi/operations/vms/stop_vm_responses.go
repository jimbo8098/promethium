// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// StopVMOKCode is the HTTP code returned for type StopVMOK
const StopVMOKCode int = 200

/*StopVMOK successful operation

swagger:response stopVmOK
*/
type StopVMOK struct {

	/*
	  In: Body
	*/
	Payload *models.VM `json:"body,omitempty"`
}

// NewStopVMOK creates StopVMOK with default headers values
func NewStopVMOK() *StopVMOK {

	return &StopVMOK{}
}

// WithPayload adds the payload to the stop Vm o k response
func (o *StopVMOK) WithPayload(payload *models.VM) *StopVMOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop Vm o k response
func (o *StopVMOK) SetPayload(payload *models.VM) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopVMOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopVMBadRequestCode is the HTTP code returned for type StopVMBadRequest
const StopVMBadRequestCode int = 400

/*StopVMBadRequest Invalid ID supplied

swagger:response stopVmBadRequest
*/
type StopVMBadRequest struct {
}

// NewStopVMBadRequest creates StopVMBadRequest with default headers values
func NewStopVMBadRequest() *StopVMBadRequest {

	return &StopVMBadRequest{}
}

// WriteResponse to the client
func (o *StopVMBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// StopVMNotFoundCode is the HTTP code returned for type StopVMNotFound
const StopVMNotFoundCode int = 404

/*StopVMNotFound VM not found

swagger:response stopVmNotFound
*/
type StopVMNotFound struct {
}

// NewStopVMNotFound creates StopVMNotFound with default headers values
func NewStopVMNotFound() *StopVMNotFound {

	return &StopVMNotFound{}
}

// WriteResponse to the client
func (o *StopVMNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
