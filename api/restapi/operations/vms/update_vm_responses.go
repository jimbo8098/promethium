// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// UpdateVMOKCode is the HTTP code returned for type UpdateVMOK
const UpdateVMOKCode int = 200

/*UpdateVMOK successful operation

swagger:response updateVmOK
*/
type UpdateVMOK struct {

	/*
	  In: Body
	*/
	Payload *models.VM `json:"body,omitempty"`
}

// NewUpdateVMOK creates UpdateVMOK with default headers values
func NewUpdateVMOK() *UpdateVMOK {

	return &UpdateVMOK{}
}

// WithPayload adds the payload to the update Vm o k response
func (o *UpdateVMOK) WithPayload(payload *models.VM) *UpdateVMOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Vm o k response
func (o *UpdateVMOK) SetPayload(payload *models.VM) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateVMOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateVMBadRequestCode is the HTTP code returned for type UpdateVMBadRequest
const UpdateVMBadRequestCode int = 400

/*UpdateVMBadRequest Invalid ID supplied

swagger:response updateVmBadRequest
*/
type UpdateVMBadRequest struct {
}

// NewUpdateVMBadRequest creates UpdateVMBadRequest with default headers values
func NewUpdateVMBadRequest() *UpdateVMBadRequest {

	return &UpdateVMBadRequest{}
}

// WriteResponse to the client
func (o *UpdateVMBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateVMNotFoundCode is the HTTP code returned for type UpdateVMNotFound
const UpdateVMNotFoundCode int = 404

/*UpdateVMNotFound VM not found

swagger:response updateVmNotFound
*/
type UpdateVMNotFound struct {
}

// NewUpdateVMNotFound creates UpdateVMNotFound with default headers values
func NewUpdateVMNotFound() *UpdateVMNotFound {

	return &UpdateVMNotFound{}
}

// WriteResponse to the client
func (o *UpdateVMNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}