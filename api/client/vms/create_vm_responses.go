// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/768bit/promethium/api/models"
)

// CreateVMReader is a Reader for the CreateVM structure.
type CreateVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVMOK creates a CreateVMOK with default headers values
func NewCreateVMOK() *CreateVMOK {
	return &CreateVMOK{}
}

/*CreateVMOK handles this case with default header values.

successful operation
*/
type CreateVMOK struct {
	Payload *models.VM
}

func (o *CreateVMOK) Error() string {
	return fmt.Sprintf("[POST /vms][%d] createVmOK  %+v", 200, o.Payload)
}

func (o *CreateVMOK) GetPayload() *models.VM {
	return o.Payload
}

func (o *CreateVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VM)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVMBadRequest creates a CreateVMBadRequest with default headers values
func NewCreateVMBadRequest() *CreateVMBadRequest {
	return &CreateVMBadRequest{}
}

/*CreateVMBadRequest handles this case with default header values.

Invalid ID supplied
*/
type CreateVMBadRequest struct {
}

func (o *CreateVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /vms][%d] createVmBadRequest ", 400)
}

func (o *CreateVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVMNotFound creates a CreateVMNotFound with default headers values
func NewCreateVMNotFound() *CreateVMNotFound {
	return &CreateVMNotFound{}
}

/*CreateVMNotFound handles this case with default header values.

VM not found
*/
type CreateVMNotFound struct {
}

func (o *CreateVMNotFound) Error() string {
	return fmt.Sprintf("[POST /vms][%d] createVmNotFound ", 404)
}

func (o *CreateVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
