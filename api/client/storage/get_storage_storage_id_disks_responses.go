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

// GetStorageStorageIDDisksReader is a Reader for the GetStorageStorageIDDisks structure.
type GetStorageStorageIDDisksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageStorageIDDisksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageStorageIDDisksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStorageStorageIDDisksOK creates a GetStorageStorageIDDisksOK with default headers values
func NewGetStorageStorageIDDisksOK() *GetStorageStorageIDDisksOK {
	return &GetStorageStorageIDDisksOK{}
}

/*GetStorageStorageIDDisksOK handles this case with default header values.

OK
*/
type GetStorageStorageIDDisksOK struct {
	Payload []models.StorageDisk
}

func (o *GetStorageStorageIDDisksOK) Error() string {
	return fmt.Sprintf("[GET /storage/{storageID}/disks][%d] getStorageStorageIdDisksOK  %+v", 200, o.Payload)
}

func (o *GetStorageStorageIDDisksOK) GetPayload() []models.StorageDisk {
	return o.Payload
}

func (o *GetStorageStorageIDDisksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
