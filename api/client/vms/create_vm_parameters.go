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

// NewCreateVMParams creates a new CreateVMParams object
// with the default values initialized.
func NewCreateVMParams() *CreateVMParams {
	var ()
	return &CreateVMParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMParamsWithTimeout creates a new CreateVMParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVMParamsWithTimeout(timeout time.Duration) *CreateVMParams {
	var ()
	return &CreateVMParams{

		timeout: timeout,
	}
}

// NewCreateVMParamsWithContext creates a new CreateVMParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVMParamsWithContext(ctx context.Context) *CreateVMParams {
	var ()
	return &CreateVMParams{

		Context: ctx,
	}
}

// NewCreateVMParamsWithHTTPClient creates a new CreateVMParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVMParamsWithHTTPClient(client *http.Client) *CreateVMParams {
	var ()
	return &CreateVMParams{
		HTTPClient: client,
	}
}

/*CreateVMParams contains all the parameters to send to the API endpoint
for the create VM operation typically these are written to a http.Request
*/
type CreateVMParams struct {

	/*VMConfig
	  Create new VM instance

	*/
	VMConfig *models.NewVM

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create VM params
func (o *CreateVMParams) WithTimeout(timeout time.Duration) *CreateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create VM params
func (o *CreateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create VM params
func (o *CreateVMParams) WithContext(ctx context.Context) *CreateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create VM params
func (o *CreateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create VM params
func (o *CreateVMParams) WithHTTPClient(client *http.Client) *CreateVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create VM params
func (o *CreateVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVMConfig adds the vMConfig to the create VM params
func (o *CreateVMParams) WithVMConfig(vMConfig *models.NewVM) *CreateVMParams {
	o.SetVMConfig(vMConfig)
	return o
}

// SetVMConfig adds the vmConfig to the create VM params
func (o *CreateVMParams) SetVMConfig(vMConfig *models.NewVM) {
	o.VMConfig = vMConfig
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.VMConfig != nil {
		if err := r.SetBodyParam(o.VMConfig); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}