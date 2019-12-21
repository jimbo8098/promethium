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

// GetStorageListReader is a Reader for the GetStorageList structure.
type GetStorageListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStorageListOK creates a GetStorageListOK with default headers values
func NewGetStorageListOK() *GetStorageListOK {
	return &GetStorageListOK{}
}

/*GetStorageListOK handles this case with default header values.

OK
*/
type GetStorageListOK struct {
	Payload []models.Storage
}

func (o *GetStorageListOK) Error() string {
	return fmt.Sprintf("[GET /storage][%d] getStorageListOK  %+v", 200, o.Payload)
}

func (o *GetStorageListOK) GetPayload() []models.Storage {
	return o.Payload
}

func (o *GetStorageListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}