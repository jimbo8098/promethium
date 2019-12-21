// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewGetStorageListParams creates a new GetStorageListParams object
// with the default values initialized.
func NewGetStorageListParams() *GetStorageListParams {

	return &GetStorageListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStorageListParamsWithTimeout creates a new GetStorageListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStorageListParamsWithTimeout(timeout time.Duration) *GetStorageListParams {

	return &GetStorageListParams{

		timeout: timeout,
	}
}

// NewGetStorageListParamsWithContext creates a new GetStorageListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStorageListParamsWithContext(ctx context.Context) *GetStorageListParams {

	return &GetStorageListParams{

		Context: ctx,
	}
}

// NewGetStorageListParamsWithHTTPClient creates a new GetStorageListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStorageListParamsWithHTTPClient(client *http.Client) *GetStorageListParams {

	return &GetStorageListParams{
		HTTPClient: client,
	}
}

/*GetStorageListParams contains all the parameters to send to the API endpoint
for the get storage list operation typically these are written to a http.Request
*/
type GetStorageListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get storage list params
func (o *GetStorageListParams) WithTimeout(timeout time.Duration) *GetStorageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get storage list params
func (o *GetStorageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get storage list params
func (o *GetStorageListParams) WithContext(ctx context.Context) *GetStorageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get storage list params
func (o *GetStorageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get storage list params
func (o *GetStorageListParams) WithHTTPClient(client *http.Client) *GetStorageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get storage list params
func (o *GetStorageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStorageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}