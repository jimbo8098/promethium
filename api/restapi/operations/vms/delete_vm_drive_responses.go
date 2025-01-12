// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// DeleteVMDriveOKCode is the HTTP code returned for type DeleteVMDriveOK
const DeleteVMDriveOKCode int = 200

/*DeleteVMDriveOK successful operation

swagger:response deleteVmDriveOK
*/
type DeleteVMDriveOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewDeleteVMDriveOK creates DeleteVMDriveOK with default headers values
func NewDeleteVMDriveOK() *DeleteVMDriveOK {

	return &DeleteVMDriveOK{}
}

// WithPayload adds the payload to the delete Vm drive o k response
func (o *DeleteVMDriveOK) WithPayload(payload *models.Item) *DeleteVMDriveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete Vm drive o k response
func (o *DeleteVMDriveOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVMDriveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteVMDriveBadRequestCode is the HTTP code returned for type DeleteVMDriveBadRequest
const DeleteVMDriveBadRequestCode int = 400

/*DeleteVMDriveBadRequest Invalid ID supplied

swagger:response deleteVmDriveBadRequest
*/
type DeleteVMDriveBadRequest struct {
}

// NewDeleteVMDriveBadRequest creates DeleteVMDriveBadRequest with default headers values
func NewDeleteVMDriveBadRequest() *DeleteVMDriveBadRequest {

	return &DeleteVMDriveBadRequest{}
}

// WriteResponse to the client
func (o *DeleteVMDriveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteVMDriveNotFoundCode is the HTTP code returned for type DeleteVMDriveNotFound
const DeleteVMDriveNotFoundCode int = 404

/*DeleteVMDriveNotFound VM not found

swagger:response deleteVmDriveNotFound
*/
type DeleteVMDriveNotFound struct {
}

// NewDeleteVMDriveNotFound creates DeleteVMDriveNotFound with default headers values
func NewDeleteVMDriveNotFound() *DeleteVMDriveNotFound {

	return &DeleteVMDriveNotFound{}
}

// WriteResponse to the client
func (o *DeleteVMDriveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*DeleteVMDriveDefault unexpected error

swagger:response deleteVmDriveDefault
*/
type DeleteVMDriveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteVMDriveDefault creates DeleteVMDriveDefault with default headers values
func NewDeleteVMDriveDefault(code int) *DeleteVMDriveDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteVMDriveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete VM drive default response
func (o *DeleteVMDriveDefault) WithStatusCode(code int) *DeleteVMDriveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete VM drive default response
func (o *DeleteVMDriveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete VM drive default response
func (o *DeleteVMDriveDefault) WithPayload(payload *models.Error) *DeleteVMDriveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete VM drive default response
func (o *DeleteVMDriveDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVMDriveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
