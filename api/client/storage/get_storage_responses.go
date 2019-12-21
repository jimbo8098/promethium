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

// GetStorageReader is a Reader for the GetStorage structure.
type GetStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStorageOK creates a GetStorageOK with default headers values
func NewGetStorageOK() *GetStorageOK {
	return &GetStorageOK{}
}

/*GetStorageOK handles this case with default header values.

successful operation
*/
type GetStorageOK struct {
	Payload models.Storage
}

func (o *GetStorageOK) Error() string {
	return fmt.Sprintf("[GET /storage/{storageID}][%d] getStorageOK  %+v", 200, o.Payload)
}

func (o *GetStorageOK) GetPayload() models.Storage {
	return o.Payload
}

func (o *GetStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageBadRequest creates a GetStorageBadRequest with default headers values
func NewGetStorageBadRequest() *GetStorageBadRequest {
	return &GetStorageBadRequest{}
}

/*GetStorageBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetStorageBadRequest struct {
}

func (o *GetStorageBadRequest) Error() string {
	return fmt.Sprintf("[GET /storage/{storageID}][%d] getStorageBadRequest ", 400)
}

func (o *GetStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageNotFound creates a GetStorageNotFound with default headers values
func NewGetStorageNotFound() *GetStorageNotFound {
	return &GetStorageNotFound{}
}

/*GetStorageNotFound handles this case with default header values.

Storage not found
*/
type GetStorageNotFound struct {
}

func (o *GetStorageNotFound) Error() string {
	return fmt.Sprintf("[GET /storage/{storageID}][%d] getStorageNotFound ", 404)
}

func (o *GetStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}