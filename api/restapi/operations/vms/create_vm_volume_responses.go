// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// CreateVMVolumeOKCode is the HTTP code returned for type CreateVMVolumeOK
const CreateVMVolumeOKCode int = 200

/*CreateVMVolumeOK successful operation

swagger:response createVmVolumeOK
*/
type CreateVMVolumeOK struct {

	/*
	  In: Body
	*/
	Payload *models.VMVolume `json:"body,omitempty"`
}

// NewCreateVMVolumeOK creates CreateVMVolumeOK with default headers values
func NewCreateVMVolumeOK() *CreateVMVolumeOK {

	return &CreateVMVolumeOK{}
}

// WithPayload adds the payload to the create Vm volume o k response
func (o *CreateVMVolumeOK) WithPayload(payload *models.VMVolume) *CreateVMVolumeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Vm volume o k response
func (o *CreateVMVolumeOK) SetPayload(payload *models.VMVolume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVMVolumeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVMVolumeBadRequestCode is the HTTP code returned for type CreateVMVolumeBadRequest
const CreateVMVolumeBadRequestCode int = 400

/*CreateVMVolumeBadRequest Invalid ID supplied

swagger:response createVmVolumeBadRequest
*/
type CreateVMVolumeBadRequest struct {
}

// NewCreateVMVolumeBadRequest creates CreateVMVolumeBadRequest with default headers values
func NewCreateVMVolumeBadRequest() *CreateVMVolumeBadRequest {

	return &CreateVMVolumeBadRequest{}
}

// WriteResponse to the client
func (o *CreateVMVolumeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CreateVMVolumeNotFoundCode is the HTTP code returned for type CreateVMVolumeNotFound
const CreateVMVolumeNotFoundCode int = 404

/*CreateVMVolumeNotFound VM not found

swagger:response createVmVolumeNotFound
*/
type CreateVMVolumeNotFound struct {
}

// NewCreateVMVolumeNotFound creates CreateVMVolumeNotFound with default headers values
func NewCreateVMVolumeNotFound() *CreateVMVolumeNotFound {

	return &CreateVMVolumeNotFound{}
}

// WriteResponse to the client
func (o *CreateVMVolumeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*CreateVMVolumeDefault unexpected error

swagger:response createVmVolumeDefault
*/
type CreateVMVolumeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateVMVolumeDefault creates CreateVMVolumeDefault with default headers values
func NewCreateVMVolumeDefault(code int) *CreateVMVolumeDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateVMVolumeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create VM volume default response
func (o *CreateVMVolumeDefault) WithStatusCode(code int) *CreateVMVolumeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create VM volume default response
func (o *CreateVMVolumeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create VM volume default response
func (o *CreateVMVolumeDefault) WithPayload(payload *models.Error) *CreateVMVolumeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create VM volume default response
func (o *CreateVMVolumeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVMVolumeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}