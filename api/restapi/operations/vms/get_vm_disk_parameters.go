// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVMDiskParams creates a new GetVMDiskParams object
// no default values defined in spec.
func NewGetVMDiskParams() GetVMDiskParams {

	return GetVMDiskParams{}
}

// GetVMDiskParams contains all the bound params for the get VM disk operation
// typically these are obtained from a http.Request
//
// swagger:parameters getVMDisk
type GetVMDiskParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of VM Disk to use
	  Required: true
	  In: path
	*/
	DiskID string
	/*ID of VM to return
	  Required: true
	  In: path
	*/
	VMID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVMDiskParams() beforehand.
func (o *GetVMDiskParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDiskID, rhkDiskID, _ := route.Params.GetOK("diskID")
	if err := o.bindDiskID(rDiskID, rhkDiskID, route.Formats); err != nil {
		res = append(res, err)
	}

	rVMID, rhkVMID, _ := route.Params.GetOK("vmID")
	if err := o.bindVMID(rVMID, rhkVMID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDiskID binds and validates parameter DiskID from path.
func (o *GetVMDiskParams) bindDiskID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.DiskID = raw

	return nil
}

// bindVMID binds and validates parameter VMID from path.
func (o *GetVMDiskParams) bindVMID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.VMID = raw

	return nil
}
