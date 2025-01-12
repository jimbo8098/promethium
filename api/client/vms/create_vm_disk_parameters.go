// Code generated by go-swagger; DO NOT EDIT.

package vms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/768bit/promethium/api/models"
)

// NewCreateVMDiskParams creates a new CreateVMDiskParams object
// with the default values initialized.
func NewCreateVMDiskParams() *CreateVMDiskParams {
	var ()
	return &CreateVMDiskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMDiskParamsWithTimeout creates a new CreateVMDiskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVMDiskParamsWithTimeout(timeout time.Duration) *CreateVMDiskParams {
	var ()
	return &CreateVMDiskParams{

		timeout: timeout,
	}
}

// NewCreateVMDiskParamsWithContext creates a new CreateVMDiskParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVMDiskParamsWithContext(ctx context.Context) *CreateVMDiskParams {
	var ()
	return &CreateVMDiskParams{

		Context: ctx,
	}
}

// NewCreateVMDiskParamsWithHTTPClient creates a new CreateVMDiskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVMDiskParamsWithHTTPClient(client *http.Client) *CreateVMDiskParams {
	var ()
	return &CreateVMDiskParams{
		HTTPClient: client,
	}
}

/*CreateVMDiskParams contains all the parameters to send to the API endpoint
for the create VM disk operation typically these are written to a http.Request
*/
type CreateVMDiskParams struct {

	/*VMDiskConfid
	  Create new VM instance

	*/
	VMDiskConfid *models.NewVMDisk
	/*VMID
	  ID of VM to use

	*/
	VMID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create VM disk params
func (o *CreateVMDiskParams) WithTimeout(timeout time.Duration) *CreateVMDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create VM disk params
func (o *CreateVMDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create VM disk params
func (o *CreateVMDiskParams) WithContext(ctx context.Context) *CreateVMDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create VM disk params
func (o *CreateVMDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create VM disk params
func (o *CreateVMDiskParams) WithHTTPClient(client *http.Client) *CreateVMDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create VM disk params
func (o *CreateVMDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVMDiskConfid adds the vMDiskConfid to the create VM disk params
func (o *CreateVMDiskParams) WithVMDiskConfid(vMDiskConfid *models.NewVMDisk) *CreateVMDiskParams {
	o.SetVMDiskConfid(vMDiskConfid)
	return o
}

// SetVMDiskConfid adds the vmDiskConfid to the create VM disk params
func (o *CreateVMDiskParams) SetVMDiskConfid(vMDiskConfid *models.NewVMDisk) {
	o.VMDiskConfid = vMDiskConfid
}

// WithVMID adds the vMID to the create VM disk params
func (o *CreateVMDiskParams) WithVMID(vMID string) *CreateVMDiskParams {
	o.SetVMID(vMID)
	return o
}

// SetVMID adds the vmId to the create VM disk params
func (o *CreateVMDiskParams) SetVMID(vMID string) {
	o.VMID = vMID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VMDiskConfid != nil {
		if err := r.SetBodyParam(o.VMDiskConfid); err != nil {
			return err
		}
	}

	// path param vmID
	if err := r.SetPathParam("vmID", o.VMID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
