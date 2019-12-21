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

// NewGetNetworkListParams creates a new GetNetworkListParams object
// with the default values initialized.
func NewGetNetworkListParams() *GetNetworkListParams {

	return &GetNetworkListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkListParamsWithTimeout creates a new GetNetworkListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkListParamsWithTimeout(timeout time.Duration) *GetNetworkListParams {

	return &GetNetworkListParams{

		timeout: timeout,
	}
}

// NewGetNetworkListParamsWithContext creates a new GetNetworkListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkListParamsWithContext(ctx context.Context) *GetNetworkListParams {

	return &GetNetworkListParams{

		Context: ctx,
	}
}

// NewGetNetworkListParamsWithHTTPClient creates a new GetNetworkListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkListParamsWithHTTPClient(client *http.Client) *GetNetworkListParams {

	return &GetNetworkListParams{
		HTTPClient: client,
	}
}

/*GetNetworkListParams contains all the parameters to send to the API endpoint
for the get network list operation typically these are written to a http.Request
*/
type GetNetworkListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network list params
func (o *GetNetworkListParams) WithTimeout(timeout time.Duration) *GetNetworkListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network list params
func (o *GetNetworkListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network list params
func (o *GetNetworkListParams) WithContext(ctx context.Context) *GetNetworkListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network list params
func (o *GetNetworkListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network list params
func (o *GetNetworkListParams) WithHTTPClient(client *http.Client) *GetNetworkListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network list params
func (o *GetNetworkListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
