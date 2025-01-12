// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DestroyStorageReader is a Reader for the DestroyStorage structure.
type DestroyStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDestroyStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDestroyStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDestroyStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestroyStorageOK creates a DestroyStorageOK with default headers values
func NewDestroyStorageOK() *DestroyStorageOK {
	return &DestroyStorageOK{}
}

/*DestroyStorageOK handles this case with default header values.

OK
*/
type DestroyStorageOK struct {
}

func (o *DestroyStorageOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/{storageID}][%d] destroyStorageOK ", 200)
}

func (o *DestroyStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDestroyStorageBadRequest creates a DestroyStorageBadRequest with default headers values
func NewDestroyStorageBadRequest() *DestroyStorageBadRequest {
	return &DestroyStorageBadRequest{}
}

/*DestroyStorageBadRequest handles this case with default header values.

Invalid ID supplied
*/
type DestroyStorageBadRequest struct {
}

func (o *DestroyStorageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /storage/{storageID}][%d] destroyStorageBadRequest ", 400)
}

func (o *DestroyStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDestroyStorageNotFound creates a DestroyStorageNotFound with default headers values
func NewDestroyStorageNotFound() *DestroyStorageNotFound {
	return &DestroyStorageNotFound{}
}

/*DestroyStorageNotFound handles this case with default header values.

Storage not found
*/
type DestroyStorageNotFound struct {
}

func (o *DestroyStorageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /storage/{storageID}][%d] destroyStorageNotFound ", 404)
}

func (o *DestroyStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
