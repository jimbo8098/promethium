// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// GetImagesListOKCode is the HTTP code returned for type GetImagesListOK
const GetImagesListOKCode int = 200

/*GetImagesListOK Array of Images

swagger:response getImagesListOK
*/
type GetImagesListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Image `json:"body,omitempty"`
}

// NewGetImagesListOK creates GetImagesListOK with default headers values
func NewGetImagesListOK() *GetImagesListOK {

	return &GetImagesListOK{}
}

// WithPayload adds the payload to the get images list o k response
func (o *GetImagesListOK) WithPayload(payload []*models.Image) *GetImagesListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get images list o k response
func (o *GetImagesListOK) SetPayload(payload []*models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImagesListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Image, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetImagesListDefault unexpected error

swagger:response getImagesListDefault
*/
type GetImagesListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetImagesListDefault creates GetImagesListDefault with default headers values
func NewGetImagesListDefault(code int) *GetImagesListDefault {
	if code <= 0 {
		code = 500
	}

	return &GetImagesListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get images list default response
func (o *GetImagesListDefault) WithStatusCode(code int) *GetImagesListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get images list default response
func (o *GetImagesListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get images list default response
func (o *GetImagesListDefault) WithPayload(payload *models.Error) *GetImagesListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get images list default response
func (o *GetImagesListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImagesListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
