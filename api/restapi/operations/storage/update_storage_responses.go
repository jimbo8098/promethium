// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// UpdateStorageOKCode is the HTTP code returned for type UpdateStorageOK
const UpdateStorageOKCode int = 200

/*UpdateStorageOK successful operation

swagger:response updateStorageOK
*/
type UpdateStorageOK struct {

	/*
	  In: Body
	*/
	Payload models.Storage `json:"body,omitempty"`
}

// NewUpdateStorageOK creates UpdateStorageOK with default headers values
func NewUpdateStorageOK() *UpdateStorageOK {

	return &UpdateStorageOK{}
}

// WithPayload adds the payload to the update storage o k response
func (o *UpdateStorageOK) WithPayload(payload models.Storage) *UpdateStorageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update storage o k response
func (o *UpdateStorageOK) SetPayload(payload models.Storage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateStorageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// UpdateStorageBadRequestCode is the HTTP code returned for type UpdateStorageBadRequest
const UpdateStorageBadRequestCode int = 400

/*UpdateStorageBadRequest Invalid ID supplied

swagger:response updateStorageBadRequest
*/
type UpdateStorageBadRequest struct {
}

// NewUpdateStorageBadRequest creates UpdateStorageBadRequest with default headers values
func NewUpdateStorageBadRequest() *UpdateStorageBadRequest {

	return &UpdateStorageBadRequest{}
}

// WriteResponse to the client
func (o *UpdateStorageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateStorageNotFoundCode is the HTTP code returned for type UpdateStorageNotFound
const UpdateStorageNotFoundCode int = 404

/*UpdateStorageNotFound Storage not found

swagger:response updateStorageNotFound
*/
type UpdateStorageNotFound struct {
}

// NewUpdateStorageNotFound creates UpdateStorageNotFound with default headers values
func NewUpdateStorageNotFound() *UpdateStorageNotFound {

	return &UpdateStorageNotFound{}
}

// WriteResponse to the client
func (o *UpdateStorageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}