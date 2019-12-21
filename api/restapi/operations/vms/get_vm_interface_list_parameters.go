// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVMInterfaceListParams creates a new GetVMInterfaceListParams object
// with the default values initialized.
func NewGetVMInterfaceListParams() GetVMInterfaceListParams {

	var (
		// initialize parameters with default values

		limitDefault = int32(20)
		skipDefault  = int32(0)
	)

	return GetVMInterfaceListParams{
		Limit: &limitDefault,

		Skip: &skipDefault,
	}
}

// GetVMInterfaceListParams contains all the bound params for the get VM interface list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getVMInterfaceList
type GetVMInterfaceListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: 20
	*/
	Limit *int32
	/*
	  In: query
	  Default: 0
	*/
	Skip *int32
	/*ID of VM to use
	  Required: true
	  In: path
	*/
	VMID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetVMInterfaceListParams() beforehand.
func (o *GetVMInterfaceListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSkip, qhkSkip, _ := qs.GetOK("skip")
	if err := o.bindSkip(qSkip, qhkSkip, route.Formats); err != nil {
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

// bindLimit binds and validates parameter Limit from query.
func (o *GetVMInterfaceListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetVMInterfaceListParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int32", raw)
	}
	o.Limit = &value

	return nil
}

// bindSkip binds and validates parameter Skip from query.
func (o *GetVMInterfaceListParams) bindSkip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetVMInterfaceListParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("skip", "query", "int32", raw)
	}
	o.Skip = &value

	return nil
}

// bindVMID binds and validates parameter VMID from path.
func (o *GetVMInterfaceListParams) bindVMID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.VMID = raw

	return nil
}
