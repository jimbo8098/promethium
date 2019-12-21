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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVMListParams creates a new GetVMListParams object
// with the default values initialized.
func NewGetVMListParams() *GetVMListParams {
	var (
		limitDefault = int32(20)
		skipDefault  = int32(0)
	)
	return &GetVMListParams{
		Limit: &limitDefault,
		Skip:  &skipDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMListParamsWithTimeout creates a new GetVMListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVMListParamsWithTimeout(timeout time.Duration) *GetVMListParams {
	var (
		limitDefault = int32(20)
		skipDefault  = int32(0)
	)
	return &GetVMListParams{
		Limit: &limitDefault,
		Skip:  &skipDefault,

		timeout: timeout,
	}
}

// NewGetVMListParamsWithContext creates a new GetVMListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVMListParamsWithContext(ctx context.Context) *GetVMListParams {
	var (
		limitDefault = int32(20)
		skipDefault  = int32(0)
	)
	return &GetVMListParams{
		Limit: &limitDefault,
		Skip:  &skipDefault,

		Context: ctx,
	}
}

// NewGetVMListParamsWithHTTPClient creates a new GetVMListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVMListParamsWithHTTPClient(client *http.Client) *GetVMListParams {
	var (
		limitDefault = int32(20)
		skipDefault  = int32(0)
	)
	return &GetVMListParams{
		Limit:      &limitDefault,
		Skip:       &skipDefault,
		HTTPClient: client,
	}
}

/*GetVMListParams contains all the parameters to send to the API endpoint
for the get VM list operation typically these are written to a http.Request
*/
type GetVMListParams struct {

	/*Limit*/
	Limit *int32
	/*Since*/
	Since *int64
	/*Skip*/
	Skip *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get VM list params
func (o *GetVMListParams) WithTimeout(timeout time.Duration) *GetVMListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get VM list params
func (o *GetVMListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get VM list params
func (o *GetVMListParams) WithContext(ctx context.Context) *GetVMListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get VM list params
func (o *GetVMListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get VM list params
func (o *GetVMListParams) WithHTTPClient(client *http.Client) *GetVMListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get VM list params
func (o *GetVMListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get VM list params
func (o *GetVMListParams) WithLimit(limit *int32) *GetVMListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get VM list params
func (o *GetVMListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSince adds the since to the get VM list params
func (o *GetVMListParams) WithSince(since *int64) *GetVMListParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the get VM list params
func (o *GetVMListParams) SetSince(since *int64) {
	o.Since = since
}

// WithSkip adds the skip to the get VM list params
func (o *GetVMListParams) WithSkip(skip *int32) *GetVMListParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the get VM list params
func (o *GetVMListParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
