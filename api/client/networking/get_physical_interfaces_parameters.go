// Code generated by go-swagger; DO NOT EDIT.

package networking

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
)

// NewGetPhysicalInterfacesParams creates a new GetPhysicalInterfacesParams object
// with the default values initialized.
func NewGetPhysicalInterfacesParams() *GetPhysicalInterfacesParams {

	return &GetPhysicalInterfacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPhysicalInterfacesParamsWithTimeout creates a new GetPhysicalInterfacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPhysicalInterfacesParamsWithTimeout(timeout time.Duration) *GetPhysicalInterfacesParams {

	return &GetPhysicalInterfacesParams{

		timeout: timeout,
	}
}

// NewGetPhysicalInterfacesParamsWithContext creates a new GetPhysicalInterfacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPhysicalInterfacesParamsWithContext(ctx context.Context) *GetPhysicalInterfacesParams {

	return &GetPhysicalInterfacesParams{

		Context: ctx,
	}
}

// NewGetPhysicalInterfacesParamsWithHTTPClient creates a new GetPhysicalInterfacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPhysicalInterfacesParamsWithHTTPClient(client *http.Client) *GetPhysicalInterfacesParams {

	return &GetPhysicalInterfacesParams{
		HTTPClient: client,
	}
}

/*GetPhysicalInterfacesParams contains all the parameters to send to the API endpoint
for the get physical interfaces operation typically these are written to a http.Request
*/
type GetPhysicalInterfacesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get physical interfaces params
func (o *GetPhysicalInterfacesParams) WithTimeout(timeout time.Duration) *GetPhysicalInterfacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get physical interfaces params
func (o *GetPhysicalInterfacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get physical interfaces params
func (o *GetPhysicalInterfacesParams) WithContext(ctx context.Context) *GetPhysicalInterfacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get physical interfaces params
func (o *GetPhysicalInterfacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get physical interfaces params
func (o *GetPhysicalInterfacesParams) WithHTTPClient(client *http.Client) *GetPhysicalInterfacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get physical interfaces params
func (o *GetPhysicalInterfacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPhysicalInterfacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
