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

// GetVMListReader is a Reader for the GetVMList structure.
type GetVMListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVMListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVMListOK creates a GetVMListOK with default headers values
func NewGetVMListOK() *GetVMListOK {
	return &GetVMListOK{}
}

/*GetVMListOK handles this case with default header values.

Array of VMs
*/
type GetVMListOK struct {
	Payload []*models.VMListItem
}

func (o *GetVMListOK) Error() string {
	return fmt.Sprintf("[GET /vms][%d] getVmListOK  %+v", 200, o.Payload)
}

func (o *GetVMListOK) GetPayload() []*models.VMListItem {
	return o.Payload
}

func (o *GetVMListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMListDefault creates a GetVMListDefault with default headers values
func NewGetVMListDefault(code int) *GetVMListDefault {
	return &GetVMListDefault{
		_statusCode: code,
	}
}

/*GetVMListDefault handles this case with default header values.

unexpected error
*/
type GetVMListDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get VM list default response
func (o *GetVMListDefault) Code() int {
	return o._statusCode
}

func (o *GetVMListDefault) Error() string {
	return fmt.Sprintf("[GET /vms][%d] getVMList default  %+v", o._statusCode, o.Payload)
}

func (o *GetVMListDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVMListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}