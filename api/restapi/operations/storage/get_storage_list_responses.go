// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// GetStorageListOKCode is the HTTP code returned for type GetStorageListOK
const GetStorageListOKCode int = 200

/*GetStorageListOK OK

swagger:response getStorageListOK
*/
type GetStorageListOK struct {

	/*
	  In: Body
	*/
	Payload []models.Storage `json:"body,omitempty"`
}

// NewGetStorageListOK creates GetStorageListOK with default headers values
func NewGetStorageListOK() *GetStorageListOK {

	return &GetStorageListOK{}
}

// WithPayload adds the payload to the get storage list o k response
func (o *GetStorageListOK) WithPayload(payload []models.Storage) *GetStorageListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get storage list o k response
func (o *GetStorageListOK) SetPayload(payload []models.Storage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStorageListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]models.Storage, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
