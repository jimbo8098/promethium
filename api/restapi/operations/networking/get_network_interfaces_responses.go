// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/768bit/promethium/api/models"
)

// GetNetworkInterfacesOKCode is the HTTP code returned for type GetNetworkInterfacesOK
const GetNetworkInterfacesOKCode int = 200

/*GetNetworkInterfacesOK OK

swagger:response getNetworkInterfacesOK
*/
type GetNetworkInterfacesOK struct {

	/*
	  In: Body
	*/
	Payload []models.NetworkInterface `json:"body,omitempty"`
}

// NewGetNetworkInterfacesOK creates GetNetworkInterfacesOK with default headers values
func NewGetNetworkInterfacesOK() *GetNetworkInterfacesOK {

	return &GetNetworkInterfacesOK{}
}

// WithPayload adds the payload to the get network interfaces o k response
func (o *GetNetworkInterfacesOK) WithPayload(payload []models.NetworkInterface) *GetNetworkInterfacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get network interfaces o k response
func (o *GetNetworkInterfacesOK) SetPayload(payload []models.NetworkInterface) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNetworkInterfacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]models.NetworkInterface, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}