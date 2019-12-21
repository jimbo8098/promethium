// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/768bit/promethium/api/models"
)

// CreateStorageReader is a Reader for the CreateStorage structure.
type CreateStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateStorageOK creates a CreateStorageOK with default headers values
func NewCreateStorageOK() *CreateStorageOK {
	return &CreateStorageOK{}
}

/*CreateStorageOK handles this case with default header values.

successful operation
*/
type CreateStorageOK struct {
	Payload models.Storage
}

func (o *CreateStorageOK) Error() string {
	return fmt.Sprintf("[POST /storage][%d] createStorageOK  %+v", 200, o.Payload)
}

func (o *CreateStorageOK) GetPayload() models.Storage {
	return o.Payload
}

func (o *CreateStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
