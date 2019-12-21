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

// GetVMVolumeReader is a Reader for the GetVMVolume structure.
type GetVMVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMVolumeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVMVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetVMVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVMVolumeOK creates a GetVMVolumeOK with default headers values
func NewGetVMVolumeOK() *GetVMVolumeOK {
	return &GetVMVolumeOK{}
}

/*GetVMVolumeOK handles this case with default header values.

successful operation
*/
type GetVMVolumeOK struct {
	Payload *models.VMVolume
}

func (o *GetVMVolumeOK) Error() string {
	return fmt.Sprintf("[GET /vms/{vmID}/volumes/{volumeID}][%d] getVmVolumeOK  %+v", 200, o.Payload)
}

func (o *GetVMVolumeOK) GetPayload() *models.VMVolume {
	return o.Payload
}

func (o *GetVMVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMVolumeBadRequest creates a GetVMVolumeBadRequest with default headers values
func NewGetVMVolumeBadRequest() *GetVMVolumeBadRequest {
	return &GetVMVolumeBadRequest{}
}

/*GetVMVolumeBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetVMVolumeBadRequest struct {
}

func (o *GetVMVolumeBadRequest) Error() string {
	return fmt.Sprintf("[GET /vms/{vmID}/volumes/{volumeID}][%d] getVmVolumeBadRequest ", 400)
}

func (o *GetVMVolumeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVMVolumeNotFound creates a GetVMVolumeNotFound with default headers values
func NewGetVMVolumeNotFound() *GetVMVolumeNotFound {
	return &GetVMVolumeNotFound{}
}

/*GetVMVolumeNotFound handles this case with default header values.

VM not found
*/
type GetVMVolumeNotFound struct {
}

func (o *GetVMVolumeNotFound) Error() string {
	return fmt.Sprintf("[GET /vms/{vmID}/volumes/{volumeID}][%d] getVmVolumeNotFound ", 404)
}

func (o *GetVMVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVMVolumeDefault creates a GetVMVolumeDefault with default headers values
func NewGetVMVolumeDefault(code int) *GetVMVolumeDefault {
	return &GetVMVolumeDefault{
		_statusCode: code,
	}
}

/*GetVMVolumeDefault handles this case with default header values.

unexpected error
*/
type GetVMVolumeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get VM volume default response
func (o *GetVMVolumeDefault) Code() int {
	return o._statusCode
}

func (o *GetVMVolumeDefault) Error() string {
	return fmt.Sprintf("[GET /vms/{vmID}/volumes/{volumeID}][%d] getVMVolume default  %+v", o._statusCode, o.Payload)
}

func (o *GetVMVolumeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVMVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}